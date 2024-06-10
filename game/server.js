const express = require('express');
const fs = require('fs');
const path = require('path');
const { spawn } = require('child_process');

const app = express();
const port = 3000;

app.use(express.static(path.join(__dirname, 'public')));

app.get('/maze', (req, res) => {
    const size = req.query.size || '20x20'; // Default size is 20x20
    const [width, height] = size.split('x');
    
    if (isNaN(width) || isNaN(height)) {
        res.status(400).send('Invalid size parameter');
        return;
    }

    const mazeGen = spawn(path.join(__dirname, '..', 'bin', 'mazegen'), ['-size', `${width}x${height}`]);

    let mazeData = '';

    mazeGen.stdout.on('data', (data) => {
        mazeData += data.toString();
    });

    mazeGen.stderr.on('data', (data) => {
        console.error(`Error: ${data}`);
    });

    mazeGen.on('close', (code) => {
        if (code !== 0) {
            res.status(500).send('Error generating maze.');
            return;
        }

        const mazeMap = mazeData.trim().split('\n');
        res.json({ mazeMap });
    });
});

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
