import { useEffect, useState } from 'react';
import { v4 as uuidv4 } from 'uuid';
import './App.css';


function App() {
  const [data, setData] = useState({})

  const getData = () => {
    fetch('hallOfFame.json'
      , {
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        }
      }
    )
      .then(function (response) {
        // console.log(response)
        return response.json();
      })
      .then(function (myJson) {
        // console.log(myJson);
        setData(myJson)
      })
      .catch(e => console.log(e));
  }

  useEffect(() => {
    getData()
  }, [])

  return (
    <div className="App">
      <header className="App-header">
        {Object.entries(data) &&
          Object.entries(data).map(o => {
            console.log(o[1])
            return (
              <div key={o[0]}>
                <h3 key={o[0]}>{o[0]}</h3>
                {o[1] && o[1].map(e => {
                  return (
                    <p key={uuidv4()} style={{ margin: 0, display: "flex", justifyContent: "space-between" }}>
                      <span style={{ marginRight: "20px", color: "cyan" }}>{e.Name}</span>
                      <span style={{ color: "yellow" }}>{e.Score}</span>
                    </p>
                  )
                })}
              </div>
            )
          })
        }
      </header >
    </div >
  );
}

export default App;
