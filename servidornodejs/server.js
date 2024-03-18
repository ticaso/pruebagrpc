const express = require('express');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const cors = require('cors'); 

const PROTO_PATH = './aspirante.proto';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const aspiranteProto = grpc.loadPackageDefinition(packageDefinition);

const app = express();
app.use(cors()); 

const aspiranteServiceClient = new aspiranteProto.AspiranteService(
    'localhost:50051', 
    grpc.credentials.createInsecure()
);

app.get('/aspirante/:id', (req, res) => {
    console.log(`Recibida solicitud para aspirante con ID: ${req.params.id}`);
    const request = { id: parseInt(req.params.id) };
    aspiranteServiceClient.getAspirante(request, (error, response) => {
        if (!error) {
            console.log('Respuesta del servidor gRPC:', response);
            res.json(response);
        } else {
            console.error('Error en la solicitud al servidor gRPC:', error);
            res.status(500).send(error.message);
        }
    });
});


const PORT = 8080;
app.listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}`);
});
