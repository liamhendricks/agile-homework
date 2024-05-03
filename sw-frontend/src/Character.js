import { useState, useEffect } from 'react';
import Species from './Species.js'
import Starship from './Starship.js'
import Planet from './Planet.js'

const baseUrl = "http://localhost:8080/"
const characterUrl = baseUrl + "api/characters"

export default function Character({ search, page }) {
  const [data, setData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      let url = characterUrl + "?search=" + search;
      if (page != "") {
        url = url + "&page=" + page
      }
      const response = await fetch(url);
      const newData = await response.json();
      setData(newData);
    };

    fetchData();
  }, [search, page]);

  if (data) {
    console.log(data)
    if (data.code == 404) {
      return null
    }

    return (
      <>
      {data.results.map((val, i) => {
        return (
          <div key={i}>
            <h1>
              Character name: {val.name}
            </h1>
            <Planet id={val.homeworld}/>
            <ul>
            {val.species != null && val.species.map((s, ii) => {
              return (
                <Species id={s} key={ii} /> 
              );
            })}
            </ul>
            <ul>
            {val.starships != null && val.starships.map((s, ii) => {
              return (
                <Starship id={s} key={ii} /> 
              );
            })}
            </ul>
          </div>
        );
      })}
      </>
    )
  } else {
    return null;
  }
}
