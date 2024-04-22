import { Link } from "react-router-dom";
import { Form } from "react-router-dom";
export async function action({ request }) {
    const formData = await request.formData();
    // const obj = Object.fromEntries(formData);
    const url = "http://localhost:9000/login";
    await fetch(url, {
      method: "POST",
      body: formData,
    })
      .then((response) => console.log(response))
      .catch((error) => console.log(error));
}

export default function Login() {
  return (
    <>
      <h3>Conectare</h3>
      <Form method="post">
        <label>
          Email<input type="text"></input>
        </label>
        <br />
        <label>
          Password<input type="password"></input>
        </label>
        <br />
        <button type="submit">Submit</button>
      </Form>
      <Link to="/register">Nu aveți cont?</Link>
    </>
  );
}
