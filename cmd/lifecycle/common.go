package lifecycle

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/tinywell/fabtool/cmd/common"
)

const (
	chainFuncName = "chaincode"
	chainCmdDes   = "Operate a chaincode: install|instantiate|invoke|package|query|signpackage|upgrade|list."
	// UndefinedParamValue default undefined param value
	UndefinedParamValue = ""
)

var (
	chaincodeLang         string
	chaincodeCtorJSON     string
	chaincodePath         string
	chaincodeName         string
	chaincodeLabel        string
	chaincodePKGID        string
	chaincodeUsr          string // Not used
	chaincodeTarFile      string
	chaincodeQueryRaw     bool
	chaincodeQueryHex     bool
	chaincodeSequence     int64
	channelID             string
	chaincodeVersion      string
	policy                string
	escc                  string
	vscc                  string
	policyMarshalled      []byte
	transient             string
	isInit                bool
	collectionsConfigFile string
	collectionConfigBytes []byte
	peerAddresses         []string
	tlsRootCertFiles      []string
	connectionProfile     string
	waitForEvent          bool
	waitForEventTimeout   time.Duration
	instantiationPolicy   string
	user                  string
	org                   string
	output                string
)

var flags *pflag.FlagSet

func init() {
	resetFlags()
}

func resetFlags() {
	flags = &pflag.FlagSet{}

	flags.StringVarP(&chaincodeLang, "lang", "l", "golang",
		fmt.Sprintf("Language the %s is written in", chainFuncName))
	flags.StringVarP(&chaincodeCtorJSON, "ctor", "c", "{}",
		fmt.Sprintf("Constructor message for the %s in JSON format", chainFuncName))
	flags.StringVarP(&chaincodePath, "path", "p", UndefinedParamValue,
		fmt.Sprintf("Path to %s", chainFuncName))
	flags.StringVarP(&chaincodeName, "name", "n", UndefinedParamValue,
		"Name of the chaincode")
	flags.StringVarP(&chaincodeLabel, "label", "", UndefinedParamValue,
		"Label of the chaincode(mycc2.0)")
	flags.StringVarP(&chaincodePKGID, "packageID", "", UndefinedParamValue,
		"chaincode package ID")
	flags.StringVarP(&chaincodeVersion, "version", "v", UndefinedParamValue,
		"Version of the chaincode specified in install/instantiate/upgrade commands")
	flags.Int64VarP(&chaincodeSequence, "sequence", "", 0,
		"The sequence number of the chaincode definition for the channel")
	flags.StringVarP(&chaincodeUsr, "username", "u", UndefinedParamValue,
		"Username for chaincode operations when security is enabled")
	flags.StringVarP(&chaincodeTarFile, "tarfile", "", UndefinedParamValue,
		"chaincode tar file")
	flags.StringVarP(&channelID, "channelID", "C", "",
		"The channel on which this command should be executed")
	flags.StringVarP(&policy, "policy", "P", UndefinedParamValue,
		"The endorsement policy associated to this chaincode")
	flags.StringVarP(&escc, "escc", "E", UndefinedParamValue,
		"The name of the endorsement system chaincode to be used for this chaincode")
	flags.StringVarP(&vscc, "vscc", "V", UndefinedParamValue,
		"The name of the verification system chaincode to be used for this chaincode")
	flags.BoolVarP(&isInit, "isInit", "I", false,
		"Is this invocation for init (useful for supporting legacy chaincodes in the new lifecycle)")
	flags.StringVar(&collectionsConfigFile, "collections-config", UndefinedParamValue,
		"The fully qualified path to the collection JSON file including the file name")
	flags.StringArrayVarP(&peerAddresses, "peerAddresses", "", []string{UndefinedParamValue},
		"The addresses of the peers to connect to")
	flags.StringArrayVarP(&tlsRootCertFiles, "tlsRootCertFiles", "", []string{UndefinedParamValue},
		"If TLS is enabled, the paths to the TLS root cert files of the peers to connect to. The order and number of certs specified should match the --peerAddresses flag")
	flags.StringVarP(&connectionProfile, "connectionProfile", "", UndefinedParamValue,
		"Connection profile that provides the necessary connection information for the network. Note: currently only supported for providing peer connection information")
	flags.BoolVar(&waitForEvent, "waitForEvent", false,
		"Whether to wait for the event from each peer's deliver filtered service signifying that the 'invoke' transaction has been committed successfully")
	flags.DurationVar(&waitForEventTimeout, "waitForEventTimeout", 30*time.Second,
		"Time to wait for the event from each peer's deliver filtered service signifying that the 'invoke' transaction has been committed successfully")
	flags.StringVarP(&instantiationPolicy, "instantiate-policy", "i", "",
		"instantiation policy for the chaincode")
	flags.StringVarP(&user, "user", "", "Admin", "fabric user")
	flags.StringVarP(&org, "org", "", "Org1", "fabric org name")
	flags.StringVarP(&output, "output", "o", "", "output file name")
}

func attachFlags(cmd *cobra.Command, names []string) {
	cmdFlags := cmd.Flags()
	for _, name := range names {
		if flag := flags.Lookup(name); flag != nil {
			cmdFlags.AddFlag(flag)
		} else {
			//TODO:
		}
	}
}

func serviceInit() error {
	ser, err := common.SDKInit(config)
	if err != nil {
		return err
	}
	service = ser
	return nil
}
