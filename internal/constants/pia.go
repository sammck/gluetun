package constants

import (
	"net"

	"github.com/qdm12/private-internet-access-docker/internal/models"
)

const (
	PIAEncryptionPresetNormal = "normal"
	PIAEncryptionPresetStrong = "strong"
	PiaX509CRLNormal          = "MIICWDCCAUAwDQYJKoZIhvcNAQENBQAwgegxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDQTETMBEGA1UEBxMKTG9zQW5nZWxlczEgMB4GA1UEChMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBAsTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQDExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEKRMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxLzAtBgkqhkiG9w0BCQEWIHNlY3VyZUBwcml2YXRlaW50ZXJuZXRhY2Nlc3MuY29tFw0xNjA3MDgxOTAwNDZaFw0zNjA3MDMxOTAwNDZaMCYwEQIBARcMMTYwNzA4MTkwMDQ2MBECAQYXDDE2MDcwODE5MDA0NjANBgkqhkiG9w0BAQ0FAAOCAQEAQZo9X97ci8EcPYu/uK2HB152OZbeZCINmYyluLDOdcSvg6B5jI+ffKN3laDvczsG6CxmY3jNyc79XVpEYUnq4rT3FfveW1+Ralf+Vf38HdpwB8EWB4hZlQ205+21CALLvZvR8HcPxC9KEnev1mU46wkTiov0EKc+EdRxkj5yMgv0V2Reze7AP+NQ9ykvDScH4eYCsmufNpIjBLhpLE2cuZZXBLcPhuRzVoU3l7A9lvzG9mjA5YijHJGHNjlWFqyrn1CfYS6koa4TGEPngBoAziWRbDGdhEgJABHrpoaFYaL61zqyMR6jC0K2ps9qyZAN74LEBedEfK7tBOzWMwr58A=="
	PiaX509CRLStrong          = "MIIDWDCCAUAwDQYJKoZIhvcNAQENBQAwgegxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDQTETMBEGA1UEBxMKTG9zQW5nZWxlczEgMB4GA1UEChMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBAsTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQDExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEKRMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxLzAtBgkqhkiG9w0BCQEWIHNlY3VyZUBwcml2YXRlaW50ZXJuZXRhY2Nlc3MuY29tFw0xNjA3MDgxOTAwNDZaFw0zNjA3MDMxOTAwNDZaMCYwEQIBARcMMTYwNzA4MTkwMDQ2MBECAQYXDDE2MDcwODE5MDA0NjANBgkqhkiG9w0BAQ0FAAOCAgEAppFfEpGsasjB1QgJcosGpzbf2kfRhM84o2TlqY1ua+Gi5TMdKydA3LJcNTjlI9a0TYAJfeRX5IkpoglSUuHuJgXhP3nEvX10mjXDpcu/YvM8TdE5JV2+EGqZ80kFtBeOq94WcpiVKFTR4fO+VkOK9zwspFfb1cNs9rHvgJ1QMkRUF8PpLN6AkntHY0+6DnigtSaKqldqjKTDTv2OeH3nPoh80SGrt0oCOmYKfWTJGpggMGKvIdvU3vH9+EuILZKKIskt+1dwdfA5Bkz1GLmiQG7+9ZZBQUjBG9Dos4hfX/rwJ3eU8oUIm4WoTz9rb71SOEuUUjP5NPy9HNx2vx+cVvLsTF4ZDZaUztW9o9JmIURDtbeyqxuHN3prlPWB6aj73IIm2dsDQvs3XXwRIxs8NwLbJ6CyEuvEOVCskdM8rdADWx1J0lRNlOJ0Z8ieLLEmYAA834VN1SboB6wJIAPxQU3rcBhXqO9y8aa2oRMg8NxZ5gr+PnKVMqag1x0IxbIgLxtkXQvxXxQHEMSODzvcOfK/nBRBsqTj30P+R87sU8titOoxNeRnBDRNhdEy/QGAqGh62ShPpQUCJdnKRiRTjnil9hMQHevoSuFKeEMO30FQL7BZyo37GFU+q1WPCplVZgCP9hC8Rn5K2+f6KLFo5bhtowSmu+GY1yZtg+RTtsA="
	PIACertificateNormal      = "MIIFqzCCBJOgAwIBAgIJAKZ7D5Yv87qDMA0GCSqGSIb3DQEBDQUAMIHoMQswCQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExEzARBgNVBAcTCkxvc0FuZ2VsZXMxIDAeBgNVBAoTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQLExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEAxMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBCkTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMS8wLQYJKoZIhvcNAQkBFiBzZWN1cmVAcHJpdmF0ZWludGVybmV0YWNjZXNzLmNvbTAeFw0xNDA0MTcxNzM1MThaFw0zNDA0MTIxNzM1MThaMIHoMQswCQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExEzARBgNVBAcTCkxvc0FuZ2VsZXMxIDAeBgNVBAoTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQLExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UEAxMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBCkTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMS8wLQYJKoZIhvcNAQkBFiBzZWN1cmVAcHJpdmF0ZWludGVybmV0YWNjZXNzLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPXDL1L9tX6DGf36liA7UBTy5I869z0UVo3lImfOs/GSiFKPtInlesP65577nd7UNzzXlH/P/CnFPdBWlLp5ze3HRBCc/Avgr5CdMRkEsySL5GHBZsx6w2cayQ2EcRhVTwWpcdldeNO+pPr9rIgPrtXqT4SWViTQRBeGM8CDxAyTopTsobjSiYZCF9Ta1gunl0G/8Vfp+SXfYCC+ZzWvP+L1pFhPRqzQQ8k+wMZIovObK1s+nlwPaLyayzw9a8sUnvWB/5rGPdIYnQWPgoNlLN9HpSmsAcw2z8DXI9pIxbr74cb3/HSfuYGOLkRqrOk6h4RCOfuWoTrZup1uEOn+fw8CAwEAAaOCAVQwggFQMB0GA1UdDgQWBBQv63nQ/pJAt5tLy8VJcbHe22ZOsjCCAR8GA1UdIwSCARYwggESgBQv63nQ/pJAt5tLy8VJcbHe22ZOsqGB7qSB6zCB6DELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMRMwEQYDVQQHEwpMb3NBbmdlbGVzMSAwHgYDVQQKExdQcml2YXRlIEludGVybmV0IEFjY2VzczEgMB4GA1UECxMXUHJpdmF0ZSBJbnRlcm5ldCBBY2Nlc3MxIDAeBgNVBAMTF1ByaXZhdGUgSW50ZXJuZXQgQWNjZXNzMSAwHgYDVQQpExdQcml2YXRlIEludGVybmV0IEFjY2VzczEvMC0GCSqGSIb3DQEJARYgc2VjdXJlQHByaXZhdGVpbnRlcm5ldGFjY2Vzcy5jb22CCQCmew+WL/O6gzAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBDQUAA4IBAQAna5PgrtxfwTumD4+3/SYvwoD66cB8IcK//h1mCzAduU8KgUXocLx7QgJWo9lnZ8xUryXvWab2usg4fqk7FPi00bED4f4qVQFVfGfPZIH9QQ7/48bPM9RyfzImZWUCenK37pdw4Bvgoys2rHLHbGen7f28knT2j/cbMxd78tQc20TIObGjo8+ISTRclSTRBtyCGohseKYpTS9himFERpUgNtefvYHbn70mIOzfOJFTVqfrptf9jXa9N8Mpy3ayfodz1wiqdteqFXkTYoSDctgKMiZ6GdocK9nMroQipIQtpnwd4yBDWIyC6Bvlkrq5TQUtYDQ8z9v+DMO6iwyIDRiU"
	PIACertificateStrong      = "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"
)

func PIAGeoChoices() (choices []string) {
	servers := PIAServers()
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return choices
}

func PIAServers() []models.PIAServer {
	return []models.PIAServer{
		{Region: "AU Melbourne", IPs: []net.IP{{27, 50, 82, 131}, {27, 50, 82, 133}, {43, 250, 204, 83}, {43, 250, 204, 87}, {43, 250, 204, 89}, {43, 250, 204, 91}, {43, 250, 204, 93}, {43, 250, 204, 95}, {43, 250, 204, 97}, {43, 250, 204, 101}, {43, 250, 204, 103}, {43, 250, 204, 105}, {43, 250, 204, 107}, {43, 250, 204, 109}, {43, 250, 204, 111}, {43, 250, 204, 113}, {43, 250, 204, 115}, {43, 250, 204, 117}, {43, 250, 204, 123}, {43, 250, 204, 125}, {118, 127, 62, 227}, {221, 121, 139, 175}}},
		{Region: "AU Perth", IPs: []net.IP{{43, 250, 205, 59}, {43, 250, 205, 91}, {43, 250, 205, 95}}},
		{Region: "AU Sydney", IPs: []net.IP{{103, 13, 102, 113}, {103, 13, 102, 115}, {103, 13, 102, 117}, {103, 13, 102, 119}, {103, 13, 102, 121}, {103, 13, 102, 123}, {103, 13, 102, 127}, {118, 127, 60, 53}, {118, 127, 60, 61}, {221, 121, 145, 133}, {221, 121, 145, 143}, {221, 121, 146, 203}, {221, 121, 152, 215}}},
		{Region: "Austria", IPs: []net.IP{{185, 210, 219, 147}, {185, 210, 219, 154}, {185, 210, 219, 155}, {185, 210, 219, 156}, {185, 216, 34, 226}, {185, 216, 34, 227}, {185, 216, 34, 228}, {185, 216, 34, 229}, {185, 216, 34, 230}, {185, 216, 34, 232}, {185, 216, 34, 236}, {185, 216, 34, 237}, {185, 216, 34, 238}}},
		{Region: "Belgium", IPs: []net.IP{{77, 243, 191, 19}, {77, 243, 191, 20}, {77, 243, 191, 21}, {77, 243, 191, 22}, {77, 243, 191, 26}, {77, 243, 191, 27}, {185, 232, 21, 26}, {185, 232, 21, 27}, {185, 232, 21, 28}, {185, 232, 21, 29}}},
		{Region: "CA Montreal", IPs: []net.IP{{199, 229, 249, 132}, {199, 229, 249, 134}, {199, 229, 249, 137}, {199, 229, 249, 142}, {199, 229, 249, 144}, {199, 229, 249, 146}, {199, 229, 249, 149}, {199, 229, 249, 151}, {199, 229, 249, 152}, {199, 229, 249, 155}, {199, 229, 249, 159}, {199, 229, 249, 165}, {199, 229, 249, 169}, {199, 229, 249, 172}, {199, 229, 249, 176}, {199, 229, 249, 182}, {199, 229, 249, 183}, {199, 229, 249, 184}, {199, 229, 249, 185}, {199, 229, 249, 191}}},
		{Region: "CA Toronto", IPs: []net.IP{{172, 98, 67, 53}, {172, 98, 67, 71}, {172, 98, 67, 87}, {172, 98, 67, 100}}},
		{Region: "CA Vancouver", IPs: []net.IP{{107, 181, 189, 71}, {107, 181, 189, 74}, {107, 181, 189, 75}, {107, 181, 189, 85}, {107, 181, 189, 86}, {172, 83, 40, 18}, {172, 83, 40, 20}, {172, 83, 40, 23}, {172, 83, 40, 24}, {172, 83, 40, 98}, {172, 83, 40, 99}, {172, 83, 40, 100}, {172, 83, 40, 101}, {172, 83, 40, 105}, {172, 83, 40, 106}, {172, 83, 40, 109}, {172, 83, 40, 110}, {172, 83, 40, 111}, {172, 83, 40, 113}, {172, 83, 40, 115}}},
		{Region: "Czech Republic", IPs: []net.IP{{89, 238, 186, 226}, {89, 238, 186, 227}, {89, 238, 186, 228}, {89, 238, 186, 230}, {185, 216, 35, 66}, {185, 216, 35, 67}, {185, 216, 35, 68}, {185, 216, 35, 69}, {185, 216, 35, 70}, {185, 242, 6, 26}, {185, 242, 6, 27}, {185, 242, 6, 28}, {185, 242, 6, 29}, {185, 242, 6, 30}}},
		{Region: "DE Berlin", IPs: []net.IP{{185, 230, 127, 226}, {185, 230, 127, 229}, {185, 230, 127, 230}, {185, 230, 127, 233}, {185, 230, 127, 234}, {185, 230, 127, 235}, {185, 230, 127, 237}, {185, 230, 127, 239}, {185, 230, 127, 240}, {185, 230, 127, 242}, {185, 230, 127, 243}, {193, 176, 86, 122}, {193, 176, 86, 125}, {193, 176, 86, 134}, {193, 176, 86, 146}, {193, 176, 86, 150}, {193, 176, 86, 154}, {193, 176, 86, 162}, {193, 176, 86, 166}, {193, 176, 86, 174}}},
		{Region: "DE Frankfurt", IPs: []net.IP{{185, 220, 70, 133}, {185, 220, 70, 134}, {185, 220, 70, 135}, {185, 220, 70, 136}, {185, 220, 70, 137}, {185, 220, 70, 138}, {185, 220, 70, 141}, {185, 220, 70, 143}, {185, 220, 70, 144}, {185, 220, 70, 147}, {185, 220, 70, 148}, {185, 220, 70, 149}, {185, 220, 70, 150}, {185, 220, 70, 152}, {185, 220, 70, 163}, {185, 220, 70, 167}, {185, 220, 70, 170}, {185, 220, 70, 171}, {185, 220, 70, 172}, {185, 220, 70, 173}}},
		{Region: "Denmark", IPs: []net.IP{{82, 102, 20, 162}, {82, 102, 20, 164}, {82, 102, 20, 165}, {82, 102, 20, 167}, {82, 102, 20, 169}, {82, 102, 20, 170}, {82, 102, 20, 171}, {82, 102, 20, 172}, {82, 102, 20, 173}, {82, 102, 20, 174}, {82, 102, 20, 175}, {82, 102, 20, 176}, {82, 102, 20, 177}, {82, 102, 20, 178}, {82, 102, 20, 179}, {82, 102, 20, 180}, {82, 102, 20, 181}, {82, 102, 20, 183}, {82, 102, 20, 184}, {82, 102, 20, 230}}},
		{Region: "Finlan", IPs: []net.IP{{196, 244, 191, 2}, {196, 244, 191, 10}, {196, 244, 191, 18}, {196, 244, 191, 26}, {196, 244, 191, 34}, {196, 244, 191, 42}, {196, 244, 191, 50}, {196, 244, 191, 58}, {196, 244, 191, 66}, {196, 244, 191, 82}, {196, 244, 191, 90}, {196, 244, 191, 98}, {196, 244, 191, 106}, {196, 244, 191, 114}, {196, 244, 191, 122}, {196, 244, 191, 138}, {196, 244, 191, 146}}},
		{Region: "France", IPs: []net.IP{{194, 99, 106, 147}, {194, 99, 106, 148}, {194, 99, 106, 149}, {194, 187, 249, 35}, {194, 187, 249, 36}, {194, 187, 249, 37}, {194, 187, 249, 38}, {194, 187, 249, 44}, {194, 187, 249, 46}, {194, 187, 249, 50}, {194, 187, 249, 53}, {194, 187, 249, 57}, {194, 187, 249, 58}, {194, 187, 249, 60}, {194, 187, 249, 61}, {194, 187, 249, 62}, {194, 187, 249, 178}, {194, 187, 249, 183}, {194, 187, 249, 184}, {194, 187, 249, 186}}},
		{Region: "Hong Kong", IPs: []net.IP{{84, 17, 37, 23}, {84, 17, 37, 45}, {119, 81, 135, 5}, {119, 81, 135, 18}, {119, 81, 135, 26}, {119, 81, 135, 28}, {119, 81, 135, 29}, {119, 81, 135, 47}, {119, 81, 135, 51}, {119, 81, 135, 53}, {119, 81, 253, 214}, {119, 81, 253, 218}, {119, 81, 253, 229}, {119, 81, 253, 230}, {119, 81, 253, 241}, {119, 81, 253, 242}, {161, 202, 39, 202}, {161, 202, 39, 240}, {161, 202, 39, 251}, {161, 202, 44, 94}}},
		{Region: "Hungary", IPs: []net.IP{{185, 128, 26, 18}, {185, 128, 26, 19}, {185, 128, 26, 20}, {185, 128, 26, 21}, {185, 128, 26, 22}, {185, 128, 26, 23}, {185, 128, 26, 24}, {185, 189, 114, 98}}},
		{Region: "India", IPs: []net.IP{{150, 242, 12, 155}, {150, 242, 12, 171}, {150, 242, 12, 187}}},
		{Region: "Ireland", IPs: []net.IP{{23, 92, 127, 2}, {23, 92, 127, 10}, {23, 92, 127, 18}, {23, 92, 127, 34}, {23, 92, 127, 42}, {23, 92, 127, 50}, {23, 92, 127, 58}, {23, 92, 127, 66}}},
		{Region: "Israel", IPs: []net.IP{{31, 168, 172, 136}, {31, 168, 172, 142}, {31, 168, 172, 143}, {31, 168, 172, 145}, {31, 168, 172, 146}, {31, 168, 172, 147}}},
		{Region: "Italy", IPs: []net.IP{{82, 102, 21, 98}, {82, 102, 21, 210}, {82, 102, 21, 212}, {82, 102, 21, 213}, {82, 102, 21, 214}, {82, 102, 21, 215}, {82, 102, 21, 216}, {82, 102, 21, 217}, {82, 102, 21, 218}, {82, 102, 21, 219}}},
		{Region: "Japan", IPs: []net.IP{{103, 208, 220, 130}, {103, 208, 220, 131}, {103, 208, 220, 132}, {103, 208, 220, 133}, {103, 208, 220, 134}, {103, 208, 220, 135}, {103, 208, 220, 136}, {103, 208, 220, 137}, {103, 208, 220, 138}, {103, 208, 220, 139}, {103, 208, 220, 140}, {103, 208, 220, 141}, {103, 208, 220, 142}, {103, 208, 220, 143}}},
		{Region: "Luxembourg", IPs: []net.IP{{92, 223, 89, 133}, {92, 223, 89, 134}, {92, 223, 89, 135}, {92, 223, 89, 136}, {92, 223, 89, 137}, {92, 223, 89, 138}, {92, 223, 89, 140}, {92, 223, 89, 142}}},
		{Region: "Mexico", IPs: []net.IP{{169, 57, 0, 197}, {169, 57, 0, 200}, {169, 57, 0, 203}, {169, 57, 0, 205}, {169, 57, 0, 207}, {169, 57, 0, 210}, {169, 57, 0, 211}, {169, 57, 0, 212}, {169, 57, 0, 213}, {169, 57, 0, 217}, {169, 57, 0, 218}, {169, 57, 0, 221}, {169, 57, 0, 229}, {169, 57, 0, 230}, {169, 57, 0, 233}, {169, 57, 0, 236}, {169, 57, 0, 238}, {169, 57, 0, 243}, {169, 57, 0, 247}, {169, 57, 0, 248}}},
		{Region: "Netherlands", IPs: []net.IP{{46, 166, 137, 219}, {46, 166, 138, 131}, {46, 166, 186, 218}, {46, 166, 186, 238}, {46, 166, 186, 249}, {46, 166, 188, 214}, {46, 166, 188, 241}, {46, 166, 190, 215}, {85, 159, 236, 219}, {109, 201, 152, 15}, {109, 201, 152, 16}, {109, 201, 152, 25}, {109, 201, 152, 238}, {109, 201, 154, 147}}},
		{Region: "New Zealand", IPs: []net.IP{{43, 250, 207, 3}, {43, 250, 207, 7}}},
		{Region: "Norway", IPs: []net.IP{{82, 102, 27, 50}, {82, 102, 27, 51}, {82, 102, 27, 52}, {82, 102, 27, 53}, {82, 102, 27, 55}, {82, 102, 27, 56}, {82, 102, 27, 57}, {82, 102, 27, 74}, {82, 102, 27, 77}, {82, 102, 27, 78}, {82, 102, 27, 114}, {82, 102, 27, 115}, {82, 102, 27, 116}, {82, 102, 27, 117}, {82, 102, 27, 124}, {82, 102, 27, 125}, {82, 102, 27, 126}, {185, 206, 225, 222}, {185, 253, 97, 227}, {185, 253, 97, 228}}},
		{Region: "Poland", IPs: []net.IP{{185, 244, 214, 14}, {185, 244, 214, 194}, {185, 244, 214, 196}, {185, 244, 214, 197}, {185, 244, 214, 198}, {185, 244, 214, 199}, {185, 244, 214, 200}}},
		{Region: "Romania", IPs: []net.IP{{86, 105, 25, 66}, {86, 105, 25, 67}, {86, 105, 25, 68}, {86, 105, 25, 69}, {86, 105, 25, 70}, {86, 105, 25, 74}, {86, 105, 25, 75}, {86, 105, 25, 76}, {86, 105, 25, 77}, {86, 105, 25, 78}, {94, 176, 148, 34}, {185, 45, 12, 126}, {185, 210, 218, 98}, {185, 210, 218, 99}, {185, 210, 218, 100}, {185, 210, 218, 101}, {185, 210, 218, 104}, {185, 210, 218, 105}, {185, 210, 218, 108}}},
		{Region: "Singapore", IPs: []net.IP{{37, 120, 208, 66}, {37, 120, 208, 67}, {37, 120, 208, 68}, {37, 120, 208, 70}, {37, 120, 208, 72}, {37, 120, 208, 73}, {37, 120, 208, 74}, {37, 120, 208, 75}, {37, 120, 208, 76}, {37, 120, 208, 77}, {37, 120, 208, 78}, {37, 120, 208, 79}, {37, 120, 208, 81}, {37, 120, 208, 82}, {37, 120, 208, 83}}},
		{Region: "Spain", IPs: []net.IP{{185, 230, 124, 53}, {194, 99, 104, 26}, {194, 99, 104, 30}}},
		{Region: "Sweden", IPs: []net.IP{{45, 12, 220, 182}, {45, 12, 220, 183}, {45, 12, 220, 194}, {45, 12, 220, 195}, {45, 12, 220, 203}, {45, 12, 220, 206}, {45, 12, 220, 209}, {45, 12, 220, 217}, {45, 12, 220, 228}, {45, 12, 220, 234}, {45, 12, 220, 240}, {45, 12, 220, 243}, {45, 12, 220, 245}, {45, 12, 220, 248}, {45, 12, 220, 250}, {45, 12, 220, 253}}},
		{Region: "Switzerland", IPs: []net.IP{{82, 102, 24, 174}, {82, 102, 24, 250}, {91, 132, 136, 42}, {91, 132, 136, 43}, {91, 132, 136, 52}, {185, 156, 175, 83}, {185, 156, 175, 84}, {185, 156, 175, 85}, {185, 156, 175, 91}, {185, 156, 175, 92}, {185, 156, 175, 93}, {185, 230, 125, 36}, {185, 230, 125, 46}, {185, 230, 125, 48}, {185, 230, 125, 50}, {185, 230, 125, 52}, {185, 230, 125, 86}, {185, 230, 125, 90}, {195, 206, 105, 210}, {212, 102, 36, 1}}},
		{Region: "UAE", IPs: []net.IP{{45, 9, 250, 42}, {45, 9, 250, 46}, {45, 9, 250, 62}}},
		{Region: "UK London", IPs: []net.IP{{89, 238, 150, 7}, {89, 238, 150, 13}, {89, 238, 150, 18}, {89, 238, 154, 18}, {89, 238, 154, 19}, {89, 238, 154, 20}, {89, 238, 154, 24}, {89, 238, 154, 115}, {89, 238, 154, 118}, {89, 238, 154, 120}, {89, 238, 154, 165}, {89, 238, 154, 166}, {89, 238, 154, 171}, {89, 238, 154, 233}, {89, 238, 154, 238}, {89, 238, 154, 243}, {89, 238, 154, 245}, {89, 238, 154, 250}, {89, 238, 154, 251}, {89, 238, 154, 254}}},
		{Region: "UK Manchester", IPs: []net.IP{{89, 238, 137, 36}, {89, 238, 137, 37}, {89, 238, 137, 39}, {89, 238, 137, 40}, {89, 238, 137, 42}, {89, 238, 139, 4}, {89, 238, 139, 5}, {89, 238, 139, 6}, {89, 238, 139, 8}, {89, 238, 139, 9}, {89, 238, 139, 10}, {89, 238, 139, 11}, {89, 238, 139, 12}, {89, 238, 139, 13}, {89, 238, 139, 52}, {89, 238, 139, 54}, {89, 238, 139, 55}, {89, 238, 139, 57}, {89, 238, 139, 58}, {89, 249, 67, 220}}},
		{Region: "UK Southampton", IPs: []net.IP{{31, 24, 226, 134}, {31, 24, 226, 188}, {31, 24, 226, 189}, {31, 24, 226, 202}, {31, 24, 226, 204}, {31, 24, 226, 205}, {31, 24, 226, 209}, {31, 24, 226, 220}, {31, 24, 226, 222}, {31, 24, 226, 226}, {31, 24, 226, 227}, {31, 24, 226, 228}, {31, 24, 226, 230}, {31, 24, 226, 234}, {31, 24, 226, 237}, {31, 24, 226, 240}, {31, 24, 226, 241}, {31, 24, 226, 244}, {31, 24, 226, 245}, {31, 24, 226, 246}}},
		{Region: "US Atlanta", IPs: []net.IP{{156, 146, 46, 1}, {156, 146, 46, 198}}},
		{Region: "US California", IPs: []net.IP{{91, 207, 175, 38}, {91, 207, 175, 42}, {91, 207, 175, 55}, {91, 207, 175, 60}, {91, 207, 175, 71}, {91, 207, 175, 109}, {91, 207, 175, 110}, {91, 207, 175, 118}, {91, 207, 175, 125}, {91, 207, 175, 163}, {91, 207, 175, 172}, {91, 207, 175, 182}, {91, 207, 175, 204}, {91, 207, 175, 205}, {91, 207, 175, 234}, {185, 245, 87, 170}, {185, 245, 87, 181}, {185, 245, 87, 195}, {185, 245, 87, 199}, {185, 245, 87, 220}}},
		{Region: "US Chicago", IPs: []net.IP{{156, 146, 50, 65}, {156, 146, 50, 134}, {156, 146, 50, 198}, {156, 146, 51, 11}, {212, 102, 58, 113}, {212, 102, 59, 54}}},
		{Region: "US Denver", IPs: []net.IP{{174, 128, 225, 186}, {174, 128, 226, 18}, {174, 128, 236, 98}, {174, 128, 242, 234}, {174, 128, 243, 98}, {174, 128, 243, 106}, {174, 128, 244, 66}, {174, 128, 244, 74}, {174, 128, 245, 106}, {174, 128, 245, 122}, {174, 128, 250, 18}, {174, 128, 250, 26}, {198, 148, 90, 58}, {199, 115, 97, 202}, {199, 115, 98, 154}, {199, 115, 99, 82}, {199, 115, 99, 218}, {199, 115, 101, 178}, {199, 115, 101, 186}, {199, 115, 103, 10}}},
		{Region: "US East", IPs: []net.IP{{193, 37, 253, 38}, {193, 37, 253, 86}, {194, 59, 251, 7}, {194, 59, 251, 13}, {194, 59, 251, 22}, {194, 59, 251, 48}, {194, 59, 251, 59}, {194, 59, 251, 103}, {194, 59, 251, 131}, {194, 59, 251, 133}, {194, 59, 251, 136}, {194, 59, 251, 137}, {194, 59, 251, 152}, {194, 59, 251, 169}, {194, 59, 251, 174}, {194, 59, 251, 190}, {194, 59, 251, 213}, {194, 59, 251, 219}, {194, 59, 251, 233}, {194, 59, 251, 239}}},
		{Region: "US Florida", IPs: []net.IP{{193, 37, 252, 7}, {193, 37, 252, 8}, {193, 37, 252, 9}, {193, 37, 252, 17}, {193, 37, 252, 19}, {193, 37, 252, 34}, {193, 37, 252, 39}, {193, 37, 252, 45}, {193, 37, 252, 58}, {193, 37, 252, 59}, {193, 37, 252, 67}, {193, 37, 252, 68}, {193, 37, 252, 76}, {193, 37, 252, 82}, {193, 37, 252, 106}, {193, 37, 252, 115}, {193, 37, 252, 116}, {193, 37, 252, 117}, {193, 37, 252, 125}, {193, 37, 252, 126}}},
		{Region: "US Houston", IPs: []net.IP{{74, 81, 88, 18}, {74, 81, 88, 26}, {74, 81, 88, 42}, {74, 81, 88, 58}, {74, 81, 88, 66}, {74, 81, 88, 82}, {74, 81, 88, 90}, {74, 81, 88, 114}, {74, 81, 88, 162}, {205, 251, 148, 50}, {205, 251, 148, 58}, {205, 251, 148, 98}, {205, 251, 148, 138}, {205, 251, 148, 178}, {205, 251, 150, 154}, {205, 251, 150, 162}, {205, 251, 150, 194}, {205, 251, 150, 218}, {205, 251, 150, 226}, {205, 251, 151, 26}}},
		{Region: "US Las Vegas", IPs: []net.IP{{162, 251, 236, 2}, {162, 251, 236, 3}, {162, 251, 236, 4}, {162, 251, 236, 7}, {162, 251, 236, 8}, {162, 251, 236, 9}, {199, 127, 56, 83}, {199, 127, 56, 84}, {199, 127, 56, 86}, {199, 127, 56, 87}, {199, 127, 56, 88}, {199, 127, 56, 89}, {199, 127, 56, 91}, {199, 127, 56, 114}, {199, 127, 56, 115}, {199, 127, 56, 116}, {199, 127, 56, 117}, {199, 127, 56, 118}, {199, 127, 56, 119}, {199, 127, 56, 120}}},
		{Region: "US New York City", IPs: []net.IP{{107, 182, 230, 194}, {107, 182, 231, 24}, {107, 182, 231, 30}, {107, 182, 231, 34}, {107, 182, 231, 37}, {107, 182, 231, 38}, {107, 182, 231, 51}, {209, 95, 50, 12}, {209, 95, 50, 27}, {209, 95, 50, 50}, {209, 95, 50, 65}, {209, 95, 50, 66}, {209, 95, 50, 90}, {209, 95, 50, 93}, {209, 95, 50, 103}, {209, 95, 50, 104}, {209, 95, 50, 133}, {209, 95, 50, 144}, {209, 95, 50, 146}, {209, 95, 50, 162}}},
		{Region: "US Seattle", IPs: []net.IP{{104, 200, 154, 11}, {104, 200, 154, 21}, {104, 200, 154, 22}, {104, 200, 154, 44}, {104, 200, 154, 47}, {104, 200, 154, 56}, {104, 200, 154, 59}, {104, 200, 154, 62}, {104, 200, 154, 66}, {104, 200, 154, 67}, {104, 200, 154, 70}, {104, 200, 154, 81}, {104, 200, 154, 84}, {104, 200, 154, 87}, {104, 200, 154, 90}, {104, 200, 154, 91}, {104, 200, 154, 96}, {104, 200, 154, 97}, {104, 200, 154, 98}, {104, 200, 154, 99}}},
		{Region: "US Silicon Valley", IPs: []net.IP{{199, 116, 118, 130}, {199, 116, 118, 133}, {199, 116, 118, 137}, {199, 116, 118, 145}, {199, 116, 118, 159}, {199, 116, 118, 160}, {199, 116, 118, 178}, {199, 116, 118, 181}, {199, 116, 118, 182}, {199, 116, 118, 202}, {199, 116, 118, 205}, {199, 116, 118, 209}, {199, 116, 118, 211}, {199, 116, 118, 213}, {199, 116, 118, 217}, {199, 116, 118, 218}, {199, 116, 118, 236}, {199, 116, 118, 238}, {199, 116, 118, 244}, {199, 116, 118, 252}}},
		{Region: "US Texas", IPs: []net.IP{{162, 216, 46, 33}, {162, 216, 46, 40}, {162, 216, 46, 44}, {162, 216, 46, 57}, {162, 216, 46, 58}, {162, 216, 46, 70}, {162, 216, 46, 81}, {162, 216, 46, 96}, {162, 216, 46, 101}, {162, 216, 46, 112}, {162, 216, 46, 119}, {162, 216, 46, 123}, {162, 216, 46, 129}, {162, 216, 46, 132}, {162, 216, 46, 133}, {162, 216, 46, 141}, {162, 216, 46, 148}, {162, 216, 46, 154}, {162, 216, 46, 156}, {162, 216, 46, 168}}},
		{Region: "US Washington DC", IPs: []net.IP{{70, 32, 0, 24}, {70, 32, 0, 46}, {70, 32, 0, 53}, {70, 32, 0, 55}, {70, 32, 0, 61}, {70, 32, 0, 74}, {70, 32, 0, 97}, {70, 32, 0, 103}, {70, 32, 0, 104}, {70, 32, 0, 133}, {70, 32, 0, 135}, {70, 32, 0, 140}, {70, 32, 0, 153}, {70, 32, 0, 155}, {70, 32, 0, 166}, {70, 32, 0, 167}, {70, 32, 0, 168}, {70, 32, 0, 172}, {70, 32, 0, 177}, {70, 32, 0, 179}}},
		{Region: "US West", IPs: []net.IP{{104, 200, 151, 4}, {104, 200, 151, 10}, {104, 200, 151, 11}, {104, 200, 151, 14}, {104, 200, 151, 15}, {104, 200, 151, 17}, {104, 200, 151, 18}, {104, 200, 151, 21}, {104, 200, 151, 26}, {104, 200, 151, 30}, {104, 200, 151, 34}, {104, 200, 151, 42}, {104, 200, 151, 48}, {104, 200, 151, 54}, {104, 200, 151, 55}, {104, 200, 151, 73}, {104, 200, 151, 78}, {104, 200, 151, 82}, {104, 200, 151, 84}, {104, 200, 151, 86}}},
	}
}

const (
	PIAPortForwardURL models.URL = "http://209.222.18.222:2000"
)
