import { useState, useEffect } from 'react';

const baseUrl = "http://localhost:8080/"
const speciesUrl = baseUrl + "api/species/"

export default function Species({ id }) {
  const [data, setData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(speciesUrl + id);
      const newData = await response.json();
      setData(newData);
    };

    fetchData();
  }, [id]);
  console.log("SPECIES")

  if (data) {
    console.log(data)
    return (
      <div>
        <h2>
          Species: {data.name}
        </h2>
      </div>
    )
  } else {
    return null;
  }
}
