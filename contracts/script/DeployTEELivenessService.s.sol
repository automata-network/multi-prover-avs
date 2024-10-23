pragma solidity ^0.8.12;

import "forge-std/Script.sol";
import {VmSafe} from "forge-std/Vm.sol";

import {TransparentUpgradeableProxy, ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {EmptyContract} from "./utils/EmptyContract.sol";

import "@dcap-v3-attestation/utils/SigVerifyLib.sol";
import "@dcap-v3-attestation/lib/PEMCertChainLib.sol";
import "@dcap-v3-attestation/AutomataDcapV3Attestation.sol";
import "./utils/DcapTestUtils.t.sol";
import {TEELivenessVerifier} from "../src/core/TEELivenessVerifier.sol";
import "./utils/CRLParser.s.sol";

contract DeployTEELivenessVerifier is Script, DcapTestUtils, CRLParser {
    string internal constant defaultTcbInfoPath =
        "dcap-v3-attestation/contracts/assets/0923/tcbInfo.json";
    string internal constant defaultTcbInfoDirPath =
        "dcap-v3-attestation/contracts/assets/latest/tcb_info/";
    string internal constant defaultQeIdPath =
        "dcap-v3-attestation/contracts/assets/latest/identity.json";

    function setUp() public {}

    struct Output {
        address SigVerifyLib;
        address PEMCertChainLib;
        address AutomataDcapV3Attestation;
        address TEELivenessVerifier;
        address TEELivenessVerifierImpl;
        string object;
    }

    function getOutputFilePath() private view returns (string memory) {
        string memory env = vm.envString("ENV");
        return
            string.concat(
                vm.projectRoot(),
                "/script/output/tee_deploy_output_",
                env,
                ".json"
            );
    }

    function readJson() private returns (string memory) {
        bytes32 remark = keccak256(abi.encodePacked("remark"));
        string memory output = vm.readFile(getOutputFilePath());
        string[] memory keys = vm.parseJsonKeys(output, ".");
        for (uint i = 0; i < keys.length; i++) {
            if (keccak256(abi.encodePacked(keys[i])) == remark) {
                continue;
            }
            string memory keyPath = string(abi.encodePacked(".", keys[i]));
            vm.serializeAddress(
                output,
                keys[i],
                vm.parseJsonAddress(output, keyPath)
            );
        }
        return output;
    }

    function saveJson(string memory json) private {
        string memory finalJson = vm.serializeString(
            json,
            "remark",
            "TEELivenessVerifier"
        );
        vm.writeJson(finalJson, getOutputFilePath());
    }

    function deploySigVerifyLib() public {
        vm.startBroadcast();
        SigVerifyLib sigVerifyLib = new SigVerifyLib();
        vm.stopBroadcast();

        string memory output = readJson();
        vm.serializeAddress(output, "SigVerifyLib", address(sigVerifyLib));
        saveJson(output);
    }

    function deployPEMCertChainLib() public {
        vm.startBroadcast();
        PEMCertChainLib pemCertLib = new PEMCertChainLib();
        vm.stopBroadcast();
        string memory output = readJson();
        vm.serializeAddress(output, "PEMCertChainLib", address(pemCertLib));
        saveJson(output);
    }

    function updateAttestationConfig() public {
        string memory output = readJson();
        AutomataDcapV3Attestation attestation = AutomataDcapV3Attestation(
            vm.parseJsonAddress(output, ".AutomataDcapV3Attestation")
        );
        vm.startBroadcast();

        {
            VmSafe.DirEntry[] memory files = vm.readDir(defaultTcbInfoDirPath);
            for (uint i = 0; i < files.length; i++) {
                string memory tcbInfoJson = vm.readFile(files[i].path);
                (
                    bool tcbParsedSuccess,
                    TCBInfoStruct.TCBInfo memory parsedTcbInfo
                ) = parseTcbInfoJson(tcbInfoJson);
                require(tcbParsedSuccess, "failed to parse tcb");
                string memory fmspc = parsedTcbInfo.fmspc;
                attestation.configureTcbInfoJson(fmspc, parsedTcbInfo);
            }
        }

        {
            string memory enclaveIdJson = vm.readFile(defaultQeIdPath);

            (
                bool qeIdParsedSuccess,
                EnclaveIdStruct.EnclaveId memory parsedEnclaveId
            ) = parseEnclaveIdentityJson(enclaveIdJson);
            require(qeIdParsedSuccess, "failed to parse qeID");

            attestation.configureQeIdentityJson(parsedEnclaveId);
        }
        vm.stopBroadcast();
    }

    function deployAttestation() public {
        string memory output = readJson();
        vm.startBroadcast();
        AutomataDcapV3Attestation attestation = new AutomataDcapV3Attestation(
            vm.parseJsonAddress(output, ".SigVerifyLib"),
            vm.parseJsonAddress(output, ".PEMCertChainLib")
        );
        {
            // CRLs are provided directly in the CRLParser.s.sol script in it's DER encoded form
            bytes[] memory crl = decodeCrl(samplePckCrl);
            attestation.addRevokedCertSerialNum(0, crl);
        }
        vm.stopBroadcast();

        vm.serializeAddress(
            output,
            "AutomataDcapV3Attestation",
            address(attestation)
        );
        saveJson(output);

        updateAttestationConfig();
        verifyQuote();
    }

    function verifyQuote() public {
        string memory output = readJson();
        AutomataDcapV3Attestation attestation = AutomataDcapV3Attestation(
            vm.parseJsonAddress(output, ".AutomataDcapV3Attestation")
        );

        bytes
            memory data = hex"030002000000000009000e00939a7233f79c4ca9940a0db3957f0607f28dda234595e56eaeb7ce9b681a62cd000000000e0e100fffff0100000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000e700000000000000ce040fe9ad608f90e417897d0839886cfef2cf238e37c099eb3d8745ab5296f90000000000000000000000000000000000000000000000000000000000000000de79b29d706d9f00ddbbdf03aa7142df1d7ef1562ec5d4e9dbf95192ccbb650a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a73c92f9c85fe649134cbd089aa6e5dda332a1c334640b160ef93e5a9971d262ae73a3503ad2184f91e6dc82740d75851a18494bf0ffc21420d8d75dafd10e3ca1000007a91529a8b4a235d330142f7faa68dea398c060e6b55dfe279b42feb12cf2a5046211dd4754a1d677d8bb631e6d66904648bd526c87a7b17217140cffc8431f2b1fb4d19c4b4071656cbdb8eaa942c89a359e6e84f51827247a3ac35b08d03abb52e537eae321e112bf351e1f5b9d7eeb3c3ea01e278e65cec3af7f8bb6fdec40e0e100fffff0100000000000000000000000000000000000000000000000000000000000000000000000000000000001500000000000000e700000000000000192aa50ce1c0cef03ccf89e7b5b16b0d7978f5c2b1edcf774d87702e8154d8bf00000000000000000000000000000000000000000000000000000000000000008c4f5775d796503e96137f77c68a829a0056ac8ded70140b081b094490c57bff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100090000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000066d3aaf3395111d7e0f2298bf4b31be75deaa4e205829fc512a4468b4177e67e0000000000000000000000000000000000000000000000000000000000000000a91eb85c23f448565f1f8edc2c0a64849c05d9ebe3af7f1d503c5374ce33a8329546cd1f8c7fa25859aa6fa21b46bfb8cb2cf49d35ddc71ba7016030b7e98d822000000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f0500620e00002d2d2d2d2d424547494e2043455254494649434154452d2d2d2d2d0a4d494945386a4343424a6d674177494241674956414e556f5a4d75787a767164353268495755667233414a6e6d6253574d416f4743437147534d343942414d430a4d484178496a416742674e5642414d4d47556c756447567349464e4857434251513073675547786864475a76636d306751304578476a415942674e5642416f4d0a45556c756447567349454e76636e4276636d4630615739754d5251774567594456515148444174545957353059534244624746795954454c4d416b47413155450a4341774351304578437a414a42674e5642415954416c56544d423458445449304d444d774e7a45314d4445784d6c6f5844544d784d444d774e7a45314d4445780a4d6c6f77634445694d434147413155454177775a535735305a5777675530645949464244537942445a584a3061575a70593246305a5445614d426747413155450a43677752535735305a577767513239796347397959585270623234784644415342674e564241634d43314e68626e526849454e7359584a684d517377435159440a5651514944414a445154454c4d416b474131554542684d4356564d775754415442676371686b6a4f5051494242676771686b6a4f50514d4242774e43414153440a30594d43645a65616e49706b52704c72516e78456a34305241585258353563437a6f4c512b4336786c45734a466346465955546b3851616c477a777756676e4e0a4c7469373461464248794c68354e55616666574f6f344944446a434341776f77487759445652306a42426777466f41556c5739647a62306234656c4153636e550a3944504f4156634c336c5177617759445652306642475177596a42676f46366758495a616148523063484d364c79396863476b7564484a316333526c5a484e6c0a636e5a705932567a4c6d6c75644756734c6d4e766253397a5a3367765932567964476c6d61574e6864476c76626939324d7939775932746a636d772f593245390a6347786864475a76636d306d5a57356a62325270626d63395a4756794d42304741315564446751574242534863356e4b574a694e3278684f39523875657543500a434e4b596254414f42674e56485138424166384542414d434273417744415944565230544151482f4241497741444343416a734743537147534962345451454e0a41515343416977776767496f4d42344743697147534962345451454e4151454545426870554c6259304254596e77775554523251363630776767466c42676f710a686b69472b453042445145434d4949425654415142677371686b69472b4530424451454341514942446a415142677371686b69472b45304244514543416749420a446a415142677371686b69472b4530424451454341774942417a415142677371686b69472b4530424451454342414942417a415242677371686b69472b4530420a4451454342514943415038774551594c4b6f5a496876684e41513042416759434167442f4d42414743797147534962345451454e41514948416745424d4241470a43797147534962345451454e41514949416745414d42414743797147534962345451454e4151494a416745414d42414743797147534962345451454e4151494b0a416745414d42414743797147534962345451454e4151494c416745414d42414743797147534962345451454e4151494d416745414d42414743797147534962340a5451454e4151494e416745414d42414743797147534962345451454e4151494f416745414d42414743797147534962345451454e41514950416745414d4241470a43797147534962345451454e41514951416745414d42414743797147534962345451454e415149524167454e4d42384743797147534962345451454e415149530a4242414f44674d442f2f38424141414141414141414141414d42414743697147534962345451454e41514d45416741414d42514743697147534962345451454e0a4151514542674267616741414144415042676f71686b69472b45304244514546436745424d42344743697147534962345451454e4151594545482b5767692b640a5a43486c4264547956765a63557a67775241594b4b6f5a496876684e41513042427a41324d42414743797147534962345451454e415163424151482f4d4241470a43797147534962345451454e41516343415145414d42414743797147534962345451454e41516344415145414d416f4743437147534d343942414d43413063410a4d4551434943732f764b4849486742427a3833746866314a7a42686c547631595339546779544f724a5449627a496d3541694243705242537754772f784a49710a38337a617553376a3847367448424b3342722b71597a55704154616a4a513d3d0a2d2d2d2d2d454e442043455254494649434154452d2d2d2d2d0a2d2d2d2d2d424547494e2043455254494649434154452d2d2d2d2d0a4d4949436c6a4343416a32674177494241674956414a567658633239472b487051456e4a3150517a7a674658433935554d416f4743437147534d343942414d430a4d476778476a415942674e5642414d4d45556c756447567349464e48574342536232393049454e424d526f77474159445651514b4442464a626e526c624342440a62334a7762334a6864476c76626a45554d424947413155454277774c553246756447456751327868636d4578437a414a42674e564241674d416b4e424d5173770a435159445651514745774a56557a4165467730784f4441314d6a45784d4455774d5442614677307a4d7a41314d6a45784d4455774d5442614d484178496a41670a42674e5642414d4d47556c756447567349464e4857434251513073675547786864475a76636d306751304578476a415942674e5642416f4d45556c75644756730a49454e76636e4276636d4630615739754d5251774567594456515148444174545957353059534244624746795954454c4d416b474131554543417743513045780a437a414a42674e5642415954416c56544d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a304441516344516741454e53422f377432316c58534f0a3243757a7078773734654a423732457944476757357258437478327456544c7136684b6b367a2b5569525a436e71523770734f766771466553786c6d546c4a6c0a65546d693257597a33714f42757a43427544416642674e5648534d4547444157674251695a517a575770303069664f44744a5653763141624f536347724442530a42674e5648523845537a424a4d45656752614244686b466f64485277637a6f764c324e6c636e52705a6d6c6a5958526c63793530636e567a6447566b633256790a646d6c6a5a584d75615735305a577775593239744c306c756447567355306459556d397664454e424c6d526c636a416442674e5648513445466751556c5739640a7a62306234656c4153636e553944504f4156634c336c517744675944565230504151482f42415144416745474d42494741315564457745422f7751494d4159420a4166384341514177436759494b6f5a497a6a30454177494452774177524149675873566b6930772b6936565947573355462f32327561586530594a446a3155650a6e412b546a44316169356343494359623153416d4435786b66545670766f34556f79695359787244574c6d5552344349394e4b7966504e2b0a2d2d2d2d2d454e442043455254494649434154452d2d2d2d2d0a2d2d2d2d2d424547494e2043455254494649434154452d2d2d2d2d0a4d4949436a7a4343416a53674177494241674955496d554d316c71644e496e7a6737535655723951477a6b6e42717777436759494b6f5a497a6a3045417749770a614445614d4267474131554541777752535735305a5777675530645949464a766233516751304578476a415942674e5642416f4d45556c756447567349454e760a636e4276636d4630615739754d5251774567594456515148444174545957353059534244624746795954454c4d416b47413155454341774351304578437a414a0a42674e5642415954416c56544d423458445445344d4455794d5445774e4455784d466f58445451354d54497a4d54497a4e546b314f566f77614445614d4267470a4131554541777752535735305a5777675530645949464a766233516751304578476a415942674e5642416f4d45556c756447567349454e76636e4276636d46300a615739754d5251774567594456515148444174545957353059534244624746795954454c4d416b47413155454341774351304578437a414a42674e56424159540a416c56544d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a3044415163445167414543366e45774d4449595a4f6a2f69505773437a61454b69370a314f694f534c52466857476a626e42564a66566e6b59347533496a6b4459594c304d784f346d717379596a6c42616c54565978465032734a424b357a6c4b4f420a757a43427544416642674e5648534d4547444157674251695a517a575770303069664f44744a5653763141624f5363477244425342674e5648523845537a424a0a4d45656752614244686b466f64485277637a6f764c324e6c636e52705a6d6c6a5958526c63793530636e567a6447566b63325679646d6c6a5a584d75615735300a5a577775593239744c306c756447567355306459556d397664454e424c6d526c636a416442674e564851344546675155496d554d316c71644e496e7a673753560a55723951477a6b6e4271777744675944565230504151482f42415144416745474d42494741315564457745422f7751494d4159424166384341514577436759490a4b6f5a497a6a3045417749445351417752674968414f572f35516b522b533943695344634e6f6f774c7550524c735747662f59693747535839344267775477670a41694541344a306c72486f4d732b586f356f2f7358364f39515778485241765a55474f6452513763767152586171493d0a2d2d2d2d2d454e442043455254494649434154452d2d2d2d2d0a00";
        (bool succ, ) = attestation.verifyAttestation(data);
        require(succ);
    }

    function deployProxyAdmin() public {
        string memory output = readJson();
        vm.startBroadcast();
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        vm.stopBroadcast();
        vm.serializeAddress(output, "ProxyAdmin", address(proxyAdmin));
        saveJson(output);
    }

    function deployVerifier() public {
        uint256 version = vm.envUint("VERSION");
        require(version < 255, "version overflowed");

        uint256 attestValiditySecs = vm.envUint("ATTEST_VALIDITY_SECS");

        uint256 maxBlockNumberDiff = vm.envUint("MAX_BLOCK_NUMBER_DIFF");
        string memory output = readJson();
        ProxyAdmin proxyAdmin = ProxyAdmin(
            vm.parseJsonAddress(output, ".ProxyAdmin")
        );
        address attestationAddr = vm.parseJsonAddress(
            output,
            ".AutomataDcapV3Attestation"
        );
        address verifierProxyAddr;

        vm.startBroadcast();
        TEELivenessVerifier verifierImpl = new TEELivenessVerifier();
        bytes memory initializeCall;
        if (
            vm.keyExistsJson(output, ".TEELivenessVerifierProxy") && version > 1
        ) {
            verifierProxyAddr = vm.parseJsonAddress(
                output,
                ".TEELivenessVerifierProxy"
            );
            console.log("reuse proxy");
            console.logAddress(verifierProxyAddr);
            console.logAddress(address(proxyAdmin));
        } else {
            console.log("Deploy new proxy");
            EmptyContract emptyContract = new EmptyContract();
            verifierProxyAddr = address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(proxyAdmin),
                    ""
                )
            );
        }
        if (version <= 1) {
            initializeCall = abi.encodeWithSelector(
                TEELivenessVerifier.initialize.selector,
                msg.sender,
                address(attestationAddr),
                maxBlockNumberDiff,
                attestValiditySecs
            );
        } else {
            initializeCall = abi.encodeWithSelector(
                TEELivenessVerifier.reinitialize.selector,
                version,
                msg.sender,
                address(attestationAddr),
                maxBlockNumberDiff,
                attestValiditySecs
            );
        }
        
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(verifierProxyAddr),
            address(verifierImpl),
            initializeCall
        );
        vm.stopBroadcast();

        vm.serializeAddress(
            output,
            "TEELivenessVerifierProxy",
            verifierProxyAddr
        );
        vm.serializeAddress(
            output,
            "TEELivenessVerifierImpl",
            address(verifierImpl)
        );
        saveJson(output);
    }

    function all() public {
        deploySigVerifyLib();
        deployPEMCertChainLib();
        deployAttestation();
        deployProxyAdmin();
        deployVerifier();
    }
}
