const got = require('got')
const jsdom = require("jsdom");
const { JSDOM } = jsdom;
const proxyList = require("./proxy.json")
const { HttpProxyAgent, HttpsProxyAgent } = require('hpagent');

async function monitor(SKU) {
    let available = false
    let proxyCount = 0
    while (true) {
        try {
            if (proxyCount == proxyList.length - 1) {
                proxyCount = 0
            }
            let pro = proxyList[proxyCount].split(":")
            proxyCount++
            console.log(`http://${pro[2]}:${pro[3]}@${pro[0]}:${pro[1]}`)
            const httpsAgent = new HttpsProxyAgent({
                keepAlive: true,
                keepAliveMsecs: 60000,
                maxSockets: 256,
                maxFreeSockets: 256,
                proxy: `http://${pro[2]}:${pro[3]}@${pro[0]}:${pro[1]}`,
            });
            const httpAgent = new HttpProxyAgent({
                    keepAlive: true,
                    keepAliveMsecs: 60000,
                    maxSockets: 256,
                    maxFreeSockets: 256,
                    proxy: `http://${pro[2]}:${pro[3]}@${pro[0]}:${pro[1]}`,
                });
                const clientConfig = {
                    timeout: 30000,
                    throwHttpErrors: false,
                    agent :  {
                        http: httpAgent,
                        https: httpsAgent
                    }// TODO : Keep in mind

                }
            const client1 = got.extend(clientConfig)
            console.log(`Monitoring Product : ${SKU}`)

            let options = {
                method: "GET",
                url: `https://httpbin.org/ip`,
                headers : {
                    "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "accept-language": "en-US,en;q=0.9",
                    "cache-control": "no-cache",
                    "downlink": "10",
                    "ect": "4g",
                    "pragma": "no-cache",
                    "rtt": "50",
                    "sec-fetch-dest": "document",
                    "sec-fetch-mode": "navigate",
                    "sec-fetch-site": "none",
                    "sec-fetch-user": "?1",
                    "upgrade-insecure-requests": "1"
                }
            }

            
            let monitor = await client1(options)
            console.log(monitor.body)
            const { document } = (new JSDOM(monitor?.body)).window
            let isPresent = false
            let currentAvailability
            console.log(monitor.statusCode)
            switch(monitor.statusCode) {
                case 200:
                case 201:
                    document.querySelectorAll("form").forEach(e => {
                        let list = e.getElementsByTagName('input')
                        let offerId
                       for (let i = 0; i < list.length; i++) {
                        if(list[i].name === "offeringID.1") offerId = list[i]?.value
                        if(list[i].getAttribute("aria-Label")) {
                           let seller = list[i].getAttribute("aria-Label")?.split("Add to Cart from seller ")[1]?.split(" and")[0]
                           if(seller === "Amazon.com") {
                               isPresent = true
                               console.log(seller)
                               console.log(`Offer Id Listing : ${offerId}`)
                            }
                        }
                       }
                    })
                    if(isPresent) {
                        currentAvailability = true
                    } else {
                        currentAvailability = false
                    }
        
                    if(!available && currentAvailability) {
                        console.log("In Stock")
                    } 
        
                    available = currentAvailability
                    break
                case 503:
                    console.log("Detected")
                    break
                default:
                    console.log(monitor.statusCode)
                    console.log(monitor.body)
            }
          

        } catch (error) {
        console.log(error)
        }
        }
    
}

monitor("B07BYYQD2B")