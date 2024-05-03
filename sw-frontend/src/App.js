import logo from './logo.svg';
import './App.css';
import React, { useState, useEffect } from 'react';

const baseUrl = "http://localhost:8080/"
const characterUrl = baseUrl + "api/characters"

function App() {
  const [data, setData] = useState([])

  useEffect(() => {
      const fetchData = async () => {
          const response = await fetch(characterUrl, {
            mode:'cors',
            credentials: 'omit',
          });
          const newData = await response.json()
          setData(newData)
      };

      fetchData();
  }, [])

  if (data) {
    console.log(data)
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>Star Wars data</h1>
        <p>Just enter a search term to find character's and their data by name... or don't and get them all.</p>
        <input></input>
      </header>
    </div>
  );
}

export default App;
