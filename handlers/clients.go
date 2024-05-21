package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/models"
	"github.com/tarkmote-ou/scaling-thereum-2023-poly-Id-bridge/utils"
	"gorm.io/gorm"
)

type Clients struct {
	db *gorm.DB
}

type ClientRegistrationRequest struct {
	Name        string `json:"name"`
	RedirectUrl string `json:"redirectUrl"`
}

func (c Clients) Register(ctx *gin.Context) {
	body := ClientRegistrationRequest{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Bad body", "error": err.Error()})
		return
	}
	did, key, err := utils.GenerateDID("did:polyid:polygon:amoy")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Create DID error", "error": err.Error()})
		return
	}
	client := models.Client{
		Id:          did,
		Secret:      uuid.NewString(),
		PrivateKey:  key,
		Name:        body.Name,
		RedirectUrl: body.RedirectUrl,
	}
	err = client.Save(c.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, client)

}
<a href="http://epicclicks.net/click/ssp/?id=eyJhY2NlcHRfbGFuZ3VhZ2UiOiJlbiIsImFkdl91c2VyX2lkIjowLCJhZHZlcnRpc2VtZW50X2lkIjoiMTE1NzM2NCIsImJyb3dzZXIiOiJDaHJvbWUiLCJjYW1wYWlnbl9jYXRlZ29yeSI6MCwiY2FtcGFpZ25faWQiOiIxMTU3MzY0IiwiY2l0eV9nZW9uYW1lX2lkIjoxMjc1MzM5LCJjbGlja19wcmljZSI6MC4wMDAwNjMsImNvbm5lY3Rpb25fdHlwZSI6IldpRmkiLCJjb3VudHJ5X2lzbyI6IklOIiwiZGV2aWNlX3R5cGUiOiJNb2JpbGUiLCJmb3JtYXQiOiJQT1AiLCJpX3QiOjE3MTQ3Mzc5MTcsImljb24iOiIiLCJpcCI6IjEyNC42Ni4xNzQuMjMyIiwiaXNwIjoiTWljcm9zY2FuIENvbXB1dGVycyBQcml2YXRlIExpbWl0ZWQiLCJsYW5kaW5nX2lkIjowLCJvcyI6IkFuZHJvaWQiLCJvc192ZXJzaW9uIjoiMTAiLCJwYXltZW50X21vZGVsIjoiQ1BDIiwicmVkaXJlY3RfdXJsIjoiaHR0cHM6Ly9waGdvZmkuY29tL3QvS1d5cUhvSlVWeGZMdXNLSVhJLTB4cnlRVVNlTVctckNqWWYwaVRGVUtRMDBQR2VoQ1pncEN1LWJ2SlM5LWp5UUFjc1NoWmh1YXZ6a0pnR0tjbkwydGFiTldzV29ZYkkxUDF1MVA0V256Z2Z1dmtBcXE1QnZyTHBhTDBGQkZTM3dRNnVMeVdRczEwMllGNk12V0xVTmJwRGg2LVB6Um83Wmc2U2xGZWRKZGlzM05Md1c3WXhqSlVPNDNxNFUxWFIwSEtSZ1lMNi1hUjEwR1hlT1NNZEJjTTMxeTJkSVRoZ0JLaTRjWXgzRW0zbVdFOTgwckIySEpZZjAzYW1PNllRM19QOUdUQjE4Z3FKdWFRa3NKZGRHSnR4UjZtTWp5cGlUUVllc3p1N2hpNWJVWFJMY25scFl4R0JyWmdvNlg5SnNlenR4NUlLYXdRdTdfTFBTNzZqNDhMTG93eGxSWHRuT3p2Y2JnT096b0ZzS3dlZGFQejhSZ1c4Q001ajd3dWNZeXczYzBpcVM3c1dxSXZUWUNPclh5QUNzdnc9PSIsInNvdXJjZV9pZCI6MjMwLCJzdWIxIjoiIiwic3ViMiI6IiIsInN1YjMiOiIiLCJzdWI0IjoiIiwic3Vic2NyaWJlZF9hdCI6MCwic3Vic2NyaXB0aW9uX2lkIjowLCJ0cmFmZmljX2NoYW5uZWwiOiJTU1AiLCJ1YV92ZXJzaW9uIjoxMjQsInVzZXJfYWdlbnQiOiJNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgMTA7IEspIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgTW9iaWxlIFNhZmFyaS81MzcuMzYiLCJ3ZWJfdXNlcl9pZCI6MTE3fQ%3D%3D" target="_blank">http://epicclicks.net/click/ssp/?id=eyJhY2NlcHRfbGFuZ3VhZ2UiOiJlbiIsImFkdl91c2VyX2lkIjowLCJhZHZlcnRpc2VtZW50X2lkIjoiMTE1NzM2NCIsImJyb3dzZXIiOiJDaHJvbWUiLCJjYW1wYWlnbl9jYXRlZ29yeSI6MCwiY2FtcGFpZ25faWQiOiIxMTU3MzY0IiwiY2l0eV9nZW9uYW1lX2lkIjoxMjc1MzM5LCJjbGlja19wcmljZSI6MC4wMDAwNjMsImNvbm5lY3Rpb25fdHlwZSI6IldpRmkiLCJjb3VudHJ5X2lzbyI6IklOIiwiZGV2aWNlX3R5cGUiOiJNb2JpbGUiLCJmb3JtYXQiOiJQT1AiLCJpX3QiOjE3MTQ3Mzc5MTcsImljb24iOiIiLCJpcCI6IjEyNC42Ni4xNzQuMjMyIiwiaXNwIjoiTWljcm9zY2FuIENvbXB1dGVycyBQcml2YXRlIExpbWl0ZWQiLCJsYW5kaW5nX2lkIjowLCJvcyI6IkFuZHJvaWQiLCJvc192ZXJzaW9uIjoiMTAiLCJwYXltZW50X21vZGVsIjoiQ1BDIiwicmVkaXJlY3RfdXJsIjoiaHR0cHM6Ly9waGdvZmkuY29tL3QvS1d5cUhvSlVWeGZMdXNLSVhJLTB4cnlRVVNlTVctckNqWWYwaVRGVUtRMDBQR2VoQ1pncEN1LWJ2SlM5LWp5UUFjc1NoWmh1YXZ6a0pnR0tjbkwydGFiTldzV29ZYkkxUDF1MVA0V256Z2Z1dmtBcXE1QnZyTHBhTDBGQkZTM3dRNnVMeVdRczEwMllGNk12V0xVTmJwRGg2LVB6Um83Wmc2U2xGZWRKZGlzM05Md1c3WXhqSlVPNDNxNFUxWFIwSEtSZ1lMNi1hUjEwR1hlT1NNZEJjTTMxeTJkSVRoZ0JLaTRjWXgzRW0zbVdFOTgwckIySEpZZjAzYW1PNllRM19QOUdUQjE4Z3FKdWFRa3NKZGRHSnR4UjZtTWp5cGlUUVllc3p1N2hpNWJVWFJMY25scFl4R0JyWmdvNlg5SnNlenR4NUlLYXdRdTdfTFBTNzZqNDhMTG93eGxSWHRuT3p2Y2JnT096b0ZzS3dlZGFQejhSZ1c4Q001ajd3dWNZeXczYzBpcVM3c1dxSXZUWUNPclh5QUNzdnc9PSIsInNvdXJjZV9pZCI6MjMwLCJzdWIxIjoiIiwic3ViMiI6IiIsInN1YjMiOiIiLCJzdWI0IjoiIiwic3Vic2NyaWJlZF9hdCI6MCwic3Vic2NyaXB0aW9uX2lkIjowLCJ0cmFmZmljX2NoYW5uZWwiOiJTU1AiLCJ1YV92ZXJzaW9uIjoxMjQsInVzZXJfYWdlbnQiOiJNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgMTA7IEspIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgTW9iaWxlIFNhZmFyaS81MzcuMzYiLCJ3ZWJfdXNlcl9pZCI6MTE3fQ%3D%3D</a>

http://epicclicks.net/click/ssp/?id=eyJhY2NlcHRfbGFuZ3VhZ2UiOiJlbiIsImFkdl91c2VyX2lkIjoxMTYsImFkdmVydGlzZW1lbnRfaWQiOiI3ODEiLCJicm93c2VyIjoiQ2hyb21lIiwiY2FtcGFpZ25fY2F0ZWdvcnkiOjksImNhbXBhaWduX2lkIjoiMzUxIiwiY2l0eV9nZW9uYW1lX2lkIjoxMjczMjk0LCJjbGlja19wcmljZSI6MC4wMDIsImNvbm5lY3Rpb25fdHlwZSI6IkV0aGVybmV0IiwiY291bnRyeV9pc28iOiJJTiIsImRldmljZV90eXBlIjoiTW9iaWxlIiwiZm9ybWF0IjoiSW5QYWdlIiwiaV90IjoxNzE0NzM3OTMwLCJpY29uIjoiYS9pbWcvNjUvMTE2LzM1MS9ISFNrWkU3ZHowR3Z2UTJ4QkgzYkg5dDRpTDllTHdXczkzMWE3eTNULnBuZyIsImlwIjoiNDkuMTUuMjI3LjgyIiwiaXNwIjoiVm9kYWZvbmUgSWRlYSIsImxhbmRpbmdfaWQiOjAsIm9zIjoiQW5kcm9pZCIsIm9zX3ZlcnNpb24iOiIxMCIsInBheW1lbnRfbW9kZWwiOiJDUEMiLCJyZWRpcmVjdF91cmwiOiIiLCJzb3VyY2VfaWQiOjIzMCwic3ViMSI6IjQ5LjE1LjIyNy44Mi05NC4xMzAuMjIyLjIxMjo0MTEyNC00OS4xNS4yMjcuODItOTQuMTMwLjIyMi4yMTI6NDExMjQtIiwic3ViMiI6IiIsInN1YjMiOiIiLCJzdWI0IjoiIiwic3Vic2NyaWJlZF9hdCI6MCwic3Vic2NyaXB0aW9uX2lkIjowLCJ0cmFmZmljX2NoYW5uZWwiOiJFcGljQWRzIiwidWFfdmVyc2lvbiI6MTI0LCJ1c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDEwOyBLKSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvMTI0LjAuMC4wIE1vYmlsZSBTYWZhcmkvNTM3LjM2Iiwid2ViX3VzZXJfaWQiOjExN30%3D
eyJhY2NlcHRfbGFuZ3VhZ2UiOiJlbiIsImFkdl91c2VyX2lkIjowLCJhZHZlcnRpc2VtZW50X2lkIjoiMTE2MDY3NSIsImJyb3dzZXIiOiJDaHJvbWUiLCJjYW1wYWlnbl9jYXRlZ29yeSI6MCwiY2FtcGFpZ25faWQiOiIxMTYwNjc1IiwiY2l0eV9nZW9uYW1lX2lkIjoxMjc3MzMzLCJjbGlja19wcmljZSI6MC4wMDM4NDI5OTk5OTk5OTk5OTk2LCJjb25uZWN0aW9uX3R5cGUiOiJXaUZpIiwiY291bnRyeV9pc28iOiJJTiIsImRldmljZV90eXBlIjoiTW9iaWxlIiwiZm9ybWF0IjoiUHVzaCIsImlfdCI6MTcxNDc0MjY1MSwiaWNvbiI6Imh0dHBzOi8vaW1nLmNkbi5ob3VzZS9pLzEvS3QwOUowMGlIMExwT3hyYlR3a0lSYkFKWnBCNm9VTm9CNVNCT0tOOFNNSnRLdk9GaXlEczcycVZfc3cxc3hmWmJxSUpGTXU3Z0JJSjBBWmhMZEhOSHZqZnc4OWtiZUc1ZmJFYW05U1E0RmROdHhkcEFaMWQ4SnFqa25ITmdNcnRkMFpFeXZlRm1XcHBLZ0t0Y3llaDkyTUhxNVBKMmhhYU1lX19XTVVWZVJPUC16VFlNVEh1UUdyTVdQaVc1U3VJWWZLR0I5alhzZ3VHeDNtNyIsImlwIjoiMTUyLjU4LjE5NC40OCIsImlzcCI6IkppbyIsImxhbmRpbmdfaWQiOjEsIm9zIjoiQW5kcm9pZCIsIm9zX3ZlcnNpb24iOiIxMCIsInBheW1lbnRfbW9kZWwiOiJDUEMiLCJyZWRpcmVjdF91cmwiOiJodHRwczovL3BoZ29maS5jb20vdC9mbEM2RWtfZnpDeGFlZm41aXVUeWRfZmlJeUg2Z2JGeThNNk00TmtSX2dDNkJKbFEyUkZiVUY4Sm1vQS15SkdIaVNZakZ4STEyeTBaOTBJSFpJampRS3FXcExuOVN6MldTUFg4WVNvcUx4RWF3RTZzaFQ3M2psMHBiaHBsaXlxcGs5aFVYZld6UmRnVzBFdmZBaGk1Zy1xX2R6Tm1TenlDeDVkYWwxdUxpd2M4YnZ3S25CcFZTU05nYk9Sclh3bFpvVFdwTk96cXNWOEdKTUhQQ3lMaktaa3RTZUVhd3BLUmZFX3FmMmVkdGxzeEZWeGlJb1FpM1lqdUxGaXMzZXo5YnN1VC1lRmRXU1g3LVVNMkZIVWh1Y1g5dU1LYnJrc2dDdHNWZUhoODQ3VkFlTk9JOEhLWjV1al9vYkUzUHY0eFdsY2NIaHBlbnFBX0VKYWpZa0xlb3Q5OXNDZmFzQkFxRThHWThKdXEtUjBYX1R5Nm8zSVl5M3hvLXcteGMydGFrT0s2VElOd3ViRENjUmQ2dGI4N2FWMUFhZz09Iiwic291cmNlX2lkIjoyMzAsInN1YjEiOiIiLCJzdWIyIjoiIiwic3ViMyI6IiIsInN1YjQiOiIiLCJzdWJzY3JpYmVkX2F0IjoxNzE0NTQ5NTc3LCJzdWJzY3JpcHRpb25faWQiOjEzMDM5NjYsInRyYWZmaWNfY2hhbm5lbCI6IlNTUCIsInVhX3ZlcnNpb24iOjEyMywidXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChMaW51eDsgQW5kcm9pZCAxMDsgSykgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEyMy4wLjAuMCBNb2JpbGUgU2FmYXJpLzUzNy4zNiIsIndlYl91c2VyX2lkIjoxMTd9%3D%3D