package main 

import (
	"os"
)


func TranscribeCall(call *Call) {
	// call.Audio is the audio data

	// Create POST request to <ASR_HOST>/asr with audio_file=call.Audio
	// Assumes you are using https://github.com/ahmetoner/whisper-asr-webservice for transcription
	// Transcription is disabled if ASR_HOST is not set.
	ASR_HOST := os.Getenv("ASR_HOST")
	if ASR_HOST != "" {
		// TODO: Implement the POST request to ASR_HOST with audio_file=call.Audio
	}

}