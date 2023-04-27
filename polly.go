// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[Gets text from a file and produces an MP3 file containing the synthesized speech.]
// snippet-keyword:[Amazon Polly]
// snippet-keyword:[SynthesizeSpeech function]
// snippet-keyword:[Go]
// snippet-service:[polly]
// snippet-keyword:[Code Sample]
// snippet-sourcetype:[full-example]
// snippet-sourcedate:[2018-03-16]
/*
   Copyright 2010-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at
    http://aws.amazon.com/apache2.0/
   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package modules

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/polly"
    "os"
    "io"
    "strconv"
)

type JyPolly struct {
    sess *session.Session
    client *polly.Polly
}

func (jp *JyPolly) Init() {
    jp.sess = session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    jp.client = polly.New(jp.sess)
}

func (jp *JyPolly) Sing(voice string, rate int, content string, dir string, fileName string) (err error){
    s := "<speak> <prosody rate=\"" + strconv.Itoa(rate) +"%\" volume=\"x-loud\">" + content + "</prosody> </speak>"
    input := &polly.SynthesizeSpeechInput{OutputFormat: aws.String("mp3"), Text: aws.String(s), VoiceId: aws.String(voice), TextType:aws.String(polly.TextTypeSsml)}
    output, err := jp.client.SynthesizeSpeech(input)

    if err != nil {
        return err
    }

    err = os.MkdirAll(dir, os.ModePerm)

    if err != nil {
        return err
    }

    outFile, err := os.Create(dir + fileName)

    defer outFile.Close()
    
    if err != nil {
        return err
    }


    _, err = io.Copy(outFile, output.AudioStream)
    if err != nil {
        return err
    }

    return
}