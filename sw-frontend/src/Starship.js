import { useState, useEffect } from 'react';

const baseUrl = "http://localhost:8080/"
const starshipUrl = baseUrl + "api/starships/"

export default function Starship({ id }) {
  const [data, setData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(starshipUrl + id);
      const newData = await response.json();
      setData(newData);
    };

    fetchData();
  }, [id]);
  console.log("STARSHIP")

  if (data) {
    console.log(data)
    return (
      <div>
        <h2>
          Starship: {data.name}
        </h2>
        <p>Class: {data.starship_class}</p>
        <p>Cargo capacity: {data.cargo_capacity}</p>
      </div>
    )
  } else {
    return null;
  }
}
