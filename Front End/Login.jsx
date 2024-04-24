import { useEffect, useState } from "react";
import { Link, redirect } from "react-router-dom";
import { Form } from "react-router-dom";
export async function action({ request }) {
  const formData = await request.formData();
  // var object = {};
  // formData.forEach((value, key) => (object[key] = value));
  // const obj = Object.fromEntries(formData);
  const url = "/api/login";

  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => {
      // console.log(response);
    })
    .then((response) => {
      // redirect("/wrongcredentials");
      // console.log("raspuns din login");
      // alert("No user with this email/password!");
    })
    .catch((error) => {});

  // const [ceva, setCeva] = useState("ceva");
  // if (ceva === "ceva") setCeva({ ceva: "ceva" });
  // else if (ceva == { ceva: "ceva" }) setCeva({ nustiu: "nustiu" });
  location.reload();
  return 0;
}

export default function Login() {
  return (
    <>
      <h3>Conectare</h3>
      <Form method="post">
        <label>
          Email<input type="text" name="email" required></input>
        </label>
        <br />
        <label>
          Password<input type="password" name="parola" required></input>
        </label>
        <br />
        <button type="submit">Submit</button>
      </Form>
      <Link to="/register">Nu aveți cont?</Link>
    </>
  );
}
