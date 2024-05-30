{-# LANGUAGE OverloadedStrings #-}

import Network.Wai
import qualified Network.Wai as Wai
import Network.Wai.Handler.Warp (run)
import Network.HTTP.Client
import Network.HTTP.Client.TLS
import qualified Network.HTTP.Client as HTTP
import Lib 
import qualified Data.ByteString.Char8 as BS

targetServiceUrl :: String
targetServiceUrl = "http://localhost:8090" 

proxyApp :: Manager -> Application
proxyApp manager req respond = do
    let targetUrl = targetServiceUrl ++ BS.unpack (Wai.rawPathInfo req) ++ BS.unpack (Wai.rawQueryString req)
    initReq <- parseRequest targetUrl
    body <- Wai.strictRequestBody req
    let reqToSend = initReq {
            HTTP.method = Wai.requestMethod req,
            HTTP.requestHeaders = Wai.requestHeaders req,
            HTTP.requestBody = HTTP.RequestBodyLBS body
        }
    response <- httpLbs reqToSend manager
    respond $ responseLBS (HTTP.responseStatus response) (HTTP.responseHeaders response) (HTTP.responseBody response)

main :: IO ()
main = do
    manager <- newManager tlsManagerSettings
    cb <- initState 1 5 20.0 10 75.0 10 True
    let wrappedApp = circuitBreakerMiddleware cb (proxyApp manager)
    run 8081 wrappedApp
