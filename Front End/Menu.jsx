import { Link } from "react-router-dom";
export default function Menu({ userData, setUserData }) {
  return (
    <>
      <h2>Menu, {userData["nume"]}</h2><Link to="/ceva">Link spre Ceva</Link>
    </>
  );
}
