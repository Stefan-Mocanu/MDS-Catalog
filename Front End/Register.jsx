import { redirect } from "react-router-dom";
import { Form } from "react-router-dom";

export async function action({ request }) {
  const formData = await request.formData();
  // const obj = Object.fromEntries(formData);
  const url = "http://localhost:9000/signin";
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => console.log(response))
    .catch((error) => console.log(error));
  return redirect("/login");
}

export default function Register() {
  return (
    <>
      <h2>Register</h2>
      <Form method="post">
        <label>
          Last Name
          <input type="text" name="nume" />
        </label>
        <br />
        <label>
          First Name
          <input type="text" name="prenume" />
        </label>
        <br />
        <label>
          Email
          <input type="email" name="email" />
        </label>
        <br />
        <label>
          Password
          <input type="password" name="parola" />
        </label><br />
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
