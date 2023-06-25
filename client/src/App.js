import React, { useState, useEffect } from 'react';
import { AgGridReact } from 'ag-grid-react';
// import 'ag-grid-community/dist/styles/ag-grid.css';
// import 'ag-grid-community/dist/styles/ag-theme-alpine.css';
import axios from 'axios';
import 'ag-grid-community/styles/ag-grid.css';
import 'ag-grid-community/styles/ag-theme-alpine.css';


function App() {
  const [rowData, setRowData] = useState([]);

  useEffect(() => {
    // Fonction pour récupérer les pilotes depuis l'API
    const fetchPilotes = async () => {
      try {
        const response = await axios.get('/pilotes');
        setRowData(response.data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchPilotes();
    console.log(rowData);
  }, []);

  const columnDefs = [
    { headerName: 'ID', field: 'id_pilote' },
    { headerName: 'Nom', field: 'nom' },
    { headerName: 'Prénom', field: 'prenom' },
    { headerName: 'Date de naissance', field: 'date_naissance' },
    { headerName: 'Catégorie', field: 'categorie' },
    { headerName: 'Années d\'expérience', field: 'annees_experience' },
  ];

  return (
    <div className="ag-theme-alpine" style={{ height: '400px', width: '100%' }}>
      <AgGridReact
        columnDefs={columnDefs}
        rowData={rowData}
        pagination={true}
        paginationPageSize={10}
      />
    </div>
  );
}

export default App;