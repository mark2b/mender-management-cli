package actions

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mender-management-cli/conf"
	"mender-management-cli/log"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadArtifact() {
	if err := uploadArtifact(); err == nil {
		log.Log.Info("Artifact uploaded")
	} else {
		log.Log.Error("Request failed. %s", err.Error())
	}
}

func uploadArtifact() (e error) {
	filename := UploadArtifactContext.ArtifactFilepath
	if fileInfo, err := os.Stat(filename); err == nil {
		if file, err := os.Open(filename); err == nil {
			if token, err := login(); err == nil {
				endpoint := conf.Config.Endpoint + "/api/management/v1/deployments/artifacts"

				var body bytes.Buffer
				writer := multipart.NewWriter(&body)
				writer.WriteField("size", fmt.Sprintf("%d", fileInfo.Size()))
				if part, err := writer.CreateFormFile("artifact", filename); err == nil {
					if _, err = io.Copy(part, file); err == nil {
						writer.Close()
						if request, err := http.NewRequest("POST", endpoint, &body); err == nil {
							request.Header.Set("Content-Type", writer.FormDataContentType())
							request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
							if _, err := httpRequest(request); err == nil {

							} else {
								e = err
							}
						} else {
							e = err
						}
					} else {
						e = err
					}
				} else {
					e = err
				}
			} else {
				e = err
			}
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func login() (token string, e error) {
	endpoint := conf.Config.Endpoint + "/api/management/v1/useradm/auth/login"
	if request, err := http.NewRequest("POST", endpoint, nil); err == nil {
		request.SetBasicAuth(conf.Config.User, conf.Config.Password)
		if responseData, err := httpRequest(request); err == nil {
			token = string(responseData)
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func httpRequest(request *http.Request) (responseData []byte, e error) {
	var httpClient = &http.Client{}
	if response, err := httpClient.Do(request); err == nil {
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			if data, err := ioutil.ReadAll(response.Body); err == nil {
				responseData = data
			} else {
				e = err
			}
		} else {
			if data, err := io.ReadAll(response.Body); err == nil {
				e = fmt.Errorf("HTTP request failed, (%d)\n%s", response.StatusCode, string(data))
			} else {
				e = fmt.Errorf("HTTP request failed, (%d)", response.StatusCode)
			}
		}
	} else {
		e = err
	}
	return
}

var UploadArtifactContext uploadArtifactContext

type uploadArtifactContext struct {
	ArtifactFilepath string
}
