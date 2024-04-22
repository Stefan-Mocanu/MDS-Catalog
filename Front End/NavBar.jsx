import { useState } from "react";
import { Link } from "react-router-dom";
export default function NavBar() {
  const [lastClickedButton, setLastClickedButton] = useState(1);

  // Function to handle button click and change color
  const changeColor = (e, buttonId) => {
    e.preventDefault();
    setLastClickedButton(buttonId);
  };
  return (
    <div id="nav_bar">
      <div id="generaloptions">
        <button>Deconnect</button>
        <button>Profile</button>
        <Link to="/addrole"><button>Add role</button></Link>
      </div>
      <div><h3>Roles:</h3></div>
      <div id="roles">
        <button>Roles</button>
      </div>
    </div>
  );
}
