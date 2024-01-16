import { useState } from "react";

import "./App.css";
import { Greet, GetOSDetails } from "../wailsjs/go/main/App";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div id="App">
      <h1 className="text-3xl font-bold underline">{count}</h1>

      <button
        className="text-3xl font-bold underline"
        onClick={() => setCount(count + 1)}
      >
        Add Count
      </button>
    </div>
  );
}

export default App;
