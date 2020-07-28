const {send} = require('micro')
const {router, get} = require('microrouter')
const fetch = require('node-fetch')

const services = {
  es: process.env.ES_SERVICE,
  fr: process.env.FR_SERVICE,
  it: process.env.IT_SERVICE
}

const i18n = async (req, res) => {
  try {
    const url = `${services[req.params.lang]}/translate?word=${req.params}`
    const response = await fetch(url)
    const data = await response.json()

    send(res, 200, data)
  }catch(err) {
    send(res, 400, { err })
  }
}

const ping = async (req, res) => await send(res, 200, { data: 'pong' })

module.exports = router(
    get('/i18n/:lang/:word', i18n),
    get('/ping', ping)
)
