import { useState, useEffect } from 'react';

const baseUrl = "http://localhost:8080/"
const planetsUrl = baseUrl + "api/planets/"

export default function Planet({ id }) {
  const [data, setData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(planetsUrl + id);
      const newData = await response.json();
      setData(newData);
    };

    fetchData();
  }, [id]);

  if (data) {
    return (
      <div>
        <h2>
          Homeworld: {data.name}
        </h2>
      </div>
    )
  } else {
    return null;
  }
}
