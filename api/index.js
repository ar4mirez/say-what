const {send} = require('micro')
const {router, get} = require('microrouter')

const i18n = async (req, res) => await send(res, 200, req.params)
const ping = async (req, res) => await send(res, 200, { data: 'pong' })

module.exports = router(
    get('/i18n/:lang/:word', i18n),
    get('/ping', ping)
)
