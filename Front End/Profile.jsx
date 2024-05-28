import { Link, useLoaderData } from "react-router-dom";

export async function loader() {
  const url = "/api/sessionActive";
  let userData = null;
  await fetch(url)
    .then((response) => response.json())
    .then((data) => {
      if (data !== false) {
        userData = { ...data };
        console.log(data);
      }
    });
  return userData;
}

export default function Profile() {
  const userData = useLoaderData();
  return (
    <>
      <h2>Profile, {userData["nume"]}</h2>
      <p>Aveti ID-ul: {userData["id"]}</p>
    </>
  );
}
