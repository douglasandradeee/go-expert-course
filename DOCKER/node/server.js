const app = require('./app')
require('dotenv').config();

const port = process.env.PORT || 3000;

app.listen(port, () => {
    console.log(`App rodando na porta: ${port}`);
});