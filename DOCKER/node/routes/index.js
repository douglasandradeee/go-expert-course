const express = require('express');
const router = express.Router();
const databaseService = require('../services/databaseService');

router.get('/', (req, res) => {
    res.send('<h1>Hello World! Maroto</h1>');
})

router.post('/people', (req, res) => {
    const { name } = req.body;

    if (!name) {
        return res.status(400).json('O campo nome é obrigatório.');
    }

    databaseService.insertPerson(name, (err, result) => {
        if (err) {
            return res.status(500).json({ error: 'Erro ao inserir pessoa no banco de dados.' });
        }
        res.status(201).json({ message: 'Pessoa inserida com sucesso.' });
    })
})

module.exports = router;