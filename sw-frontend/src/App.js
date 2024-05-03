import logo from './logo.svg';
import './App.css';
import React, { useState, useEffect } from 'react';
import Character from './Character.js';


function App() {
  const [search, setSearch] = useState("")
  const [page, setPage] = useState("")

  function handleSubmit(event) {}

  function handleChange(event) {
    setSearch(event.target.value);
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>Star Wars data</h1>
        <p>Just enter a search term to find character's and their data by name... or don't and get them all.</p>
        <form onSubmit={handleSubmit}>
          <input type="text" value={search} onChange={handleChange} />
        </form>
        <Character search={search} page={page} />
      </header>
      <footer>If I had more time I would have added pagination and made it look pretty, but you get the idea. Pls Hire me I'm cool thanks.</footer>
    </div>
  );
}

export default App;
