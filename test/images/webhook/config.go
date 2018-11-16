/*
Copyright 2017 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"crypto/tls"
)

func configTLS(_ Config) *tls.Config {
	sCert, err := tls.X509KeyPair([]byte(`-----BEGIN CERTIFICATE-----
MIIC7TCCAdWgAwIBAgIJAKy7wYXWsuV4MA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTgwNDE4MTc0NDAyWhcNMjEwMjA1MTc0NDAyWjAS
MRAwDgYDVQQDDAd3ZWJob29rMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEA4Rl/cnOYZfPuc66PIjf53taiwiJ6xELenoMsJx8fNe9NDZXE3RJmnCsGOALy
JgiaGH6XecpOXEEtuk3NJfjZP1kRDoP0RViGIk5/zgaJf+MJk9Ytc9HNVKGBDlYW
1cyrVhx2/k2ztAqiWIHtVuteHXL9fBvpreLUy1iNGYcfPlo0Xb6yvdqb/wiBHmTm
S0U7XEMpoWar9witTJaey72lYH2Xhnk3BNfdCaiRA84b9h6OSm63GQ2jNUxkJ/Th
K7/cAee4UkAT0uYEx1tGYUneiQyB8AIDl8qzix1bC9kvQKfY2pIQ3zJjSufcabCX
ZNt7wf05n6GkXV2EJcwZ5iAGJwIDAQABoxMwETAPBgNVHREECDAGhwR/AAABMA0G
CSqGSIb3DQEBCwUAA4IBAQBYUtqO/QGgRSlONS6R2pqnLG7NBA08uTCVsQuiQQcA
JYr8zAWVlxpdYyHKuCkTG0HX3rWF1trkl/Wz3UCs6dAz+RJHjGGOL80ABPjZJvkz
cNv+4Dtf6fClJsCt7+WbAQ/vBAxLp4XAbgmo8fKr5KbBNq3oNk2AVTjXivU9hyj2
zbZy4t2gbhmLBBvoeE7zNZjFp5SId8LFF3kIJIdLUJiqqnXVaYFjSjzkKWN2qOiq
lUnNlw/wHcJHQhHyMucQzgIaxEwNXXNKcOIjM0pxCruDzSW2L9Q8GOQfzsB/7B9t
fQF7S5c2+a7SxlnOifgu2i3f30BJb0y6FkQI7EZtcSUy
-----END CERTIFICATE-----`), []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4Rl/cnOYZfPuc66PIjf53taiwiJ6xELenoMsJx8fNe9NDZXE
3RJmnCsGOALyJgiaGH6XecpOXEEtuk3NJfjZP1kRDoP0RViGIk5/zgaJf+MJk9Yt
c9HNVKGBDlYW1cyrVhx2/k2ztAqiWIHtVuteHXL9fBvpreLUy1iNGYcfPlo0Xb6y
vdqb/wiBHmTmS0U7XEMpoWar9witTJaey72lYH2Xhnk3BNfdCaiRA84b9h6OSm63
GQ2jNUxkJ/ThK7/cAee4UkAT0uYEx1tGYUneiQyB8AIDl8qzix1bC9kvQKfY2pIQ
3zJjSufcabCXZNt7wf05n6GkXV2EJcwZ5iAGJwIDAQABAoIBAQDNMvBChmwT7MWg
AS4MFJkc8e7XKJVfmlHUkGFMnItYIHjOfIxEt4SAspvtrYnO8CKBV82AQdMjOGuW
HVx5VBu/KtNotxOTr7o/Re4uAzmPMonFkgZHQad7kerTBdm1NumuCy7SbkT2SIds
Raz1goR+Nhgr7kocsN4pLOUbs1jg/aZIhjjnulYSSlJZS4+35KLy4nhaTQgnoqik
/Wix0ExYMFAnLwmsOUxHQV/VMoWk9jWs7ahkLgE2CFvVgm4+N6cZDDbvHDk+5yRc
XibicQVqLMrdFHj5IXL8kIddkeDulel5TXgWzM1xLJJkxiCgWEaLQIZJeeYohLT+
aoln5ZQBAoGBAPdO0SNtmXaxZmnrQr1mydE82lqNOfu8w40q2Tq+TwOwB274YeoB
oEgGA3v/0X64XpDAbcR8mo5cfbryjnBX/RBaMe+KzIW+VIZ/3cjuEUbEVPYMkEBA
1otTMuCXnism30xiMFaX5w2RafEALOF1vB6pnQ0cMRbBqnGFmeIlSignAoGBAOkC
2/Npa1u6mXzjo1t6wZl9yhucOgNKNgqBvb5Zidc+5alpfNBHXcA7B71y335wO3/E
0qAQsAaOYwn5N0W0SbMRptAAlB+5Ly9pSLF2oeF2HJZJ61H+dkM/juYwBlnYpb/z
Fn1U3HIwqIKf8ePDvuQjwaN+4BS52Xf4Y2BVt/IBAoGBAJgAuKi234FFjjYB5LZ3
LApQBcFsVjw1DFiDApuJhxU0J418Wuoyb6p1D8UyOjhR58W+kHkZQQHJNXonRYcl
faSEW3bo78YwctFsXAv4z2OYnsPQewUTFQrzay0B47SQIuVW4HEI0nnTa7M2MV+u
Np7+D0qUjlN3W2SFAk0uMEM3AoGAIDKk8h2/GA3Q59EM4bc0yWD4bJhJ6+p1TT5g
Wc1FntiyI5bQCHfUHJwLlcIp3+7iSeWItVWY/U7voJEvchJXnMbzpgpubXPJcWO7
B8q808reaGYOfmYLMX+231gDiKbHQJ72nJr9W0od/u7bHf4OgrfuKgl+LV8BDfLk
yIsPwAECgYALaWKiORDRHGYr/YRw1lIM9AQCLsfMcIcf1y6t8ycJUwSxmUI9rRKw
w9AXrOMFioznbbfOjz4esMqWCyii/rJN++nJjiVR+PHSntebpLxe48gfYkhEMKHE
jvU19/jm2qFM69bqrpDjA3W2K4tKnesohnDoX08mrbGIEBtQCSaevQ==
-----END RSA PRIVATE KEY-----`))
	if err != nil {
		return nil
	}
	return &tls.Config{
		Certificates: []tls.Certificate{sCert},
		// TODO: uses mutual tls after we agree on what cert the apiserver should use.
		// ClientAuth:   tls.RequireAndVerifyClientCert,
	}
}
