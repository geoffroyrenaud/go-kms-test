
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"encoding/base64"
	"flag"
	"fmt"
    "os"
)


/*
KMS d'exemple
{
	"AliasName": "alias/webportal",
	"AliasArn": "arn:aws:kms:eu-west-1:625756642822:alias/webportal",
	"TargetKeyId": "9269c99d-5b42-4f47-9c89-e9932820ab4b"
}
*/

func usage()  {
	fmt.Println("Type -h for help")
	os.Exit(1)
}

func myencrypt(keyid string, toEncryptMessage string) string {
	fmt.Println("encrpyt", toEncryptMessage, "with", keyid)

	svc := kms.New(session.New())
	input := &kms.EncryptInput{
		KeyId:     aws.String(keyid),
		Plaintext: []byte(toEncryptMessage),
	}

	result, err := svc.Encrypt(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
			case kms.ErrCodeDisabledException:
				fmt.Println(kms.ErrCodeDisabledException, aerr.Error())
			case kms.ErrCodeKeyUnavailableException:
				fmt.Println(kms.ErrCodeKeyUnavailableException, aerr.Error())
			case kms.ErrCodeDependencyTimeoutException:
				fmt.Println(kms.ErrCodeDependencyTimeoutException, aerr.Error())
			case kms.ErrCodeInvalidKeyUsageException:
				fmt.Println(kms.ErrCodeInvalidKeyUsageException, aerr.Error())
			case kms.ErrCodeInvalidGrantTokenException:
				fmt.Println(kms.ErrCodeInvalidGrantTokenException, aerr.Error())
			case kms.ErrCodeInternalException:
				fmt.Println(kms.ErrCodeInternalException, aerr.Error())
			case kms.ErrCodeInvalidStateException:
				fmt.Println(kms.ErrCodeInvalidStateException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return
	}

	//s := string(result.CiphertextBlob)
	encoded := base64.StdEncoding.EncodeToString(result.CiphertextBlob)
	return encoded
}

func mydecrypt(keyid string, encryptedMessage string) string {
	fmt.Println("decrpyt", encryptedMessage, "with", keyid)

	decoded, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		fmt.Println("decode error:", err)
		return "ERROR base64decode"
	}

	svc := kms.New(session.New())
	input := &kms.DecryptInput{
		CiphertextBlob: []byte(decoded),
		KeyId: aws.String(keyid),
	}

	result, err := svc.Decrypt(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
			case kms.ErrCodeDisabledException:
				fmt.Println(kms.ErrCodeDisabledException, aerr.Error())
			case kms.ErrCodeInvalidCiphertextException:
				fmt.Println(kms.ErrCodeInvalidCiphertextException, aerr.Error())
			case kms.ErrCodeKeyUnavailableException:
				fmt.Println(kms.ErrCodeKeyUnavailableException, aerr.Error())
			case kms.ErrCodeIncorrectKeyException:
				fmt.Println(kms.ErrCodeIncorrectKeyException, aerr.Error())
			case kms.ErrCodeInvalidKeyUsageException:
				fmt.Println(kms.ErrCodeInvalidKeyUsageException, aerr.Error())
			case kms.ErrCodeDependencyTimeoutException:
				fmt.Println(kms.ErrCodeDependencyTimeoutException, aerr.Error())
			case kms.ErrCodeInvalidGrantTokenException:
				fmt.Println(kms.ErrCodeInvalidGrantTokenException, aerr.Error())
			case kms.ErrCodeInternalException:
				fmt.Println(kms.ErrCodeInternalException, aerr.Error())
			case kms.ErrCodeInvalidStateException:
				fmt.Println(kms.ErrCodeInvalidStateException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return
	}

	//fmt.Println(result)
	fmt.Println("decoded", string(result.Plaintext))
	return string(result.Plaintext)
}

func main() {

	// flag
	encryptPtr := flag.String("encrypt", "", "Encrypt message")
	decryptPtr := flag.String("decrypt", "", "Decrypt message (base64 format)")
	keyidPtr := flag.String("keyid", "", "AWS KeyID (required)")
	flag.Parse()

	if *keyidPtr == "" {
		fmt.Println("AWS KeyID is required")
		usage()
	}
	if *encryptPtr != "" {
		var encoded = myencrypt(*keyidPtr, *encryptPtr)
		fmt.Println(encoded)
	} else if *decryptPtr != "" {
		mydecrypt(*keyidPtr, *decryptPtr)
	} else {
		usage()
	} 

	
}