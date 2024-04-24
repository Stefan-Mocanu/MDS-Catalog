import { Form } from "react-router-dom";

export async function action({ request }) {
  const formData = await request.formData();
  // const obj = Object.fromEntries(formData);
  const url = "http://localhost:9000/adduser";
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => console.log(response))
    .catch((error) => console.log(error));
}

export default function AddUser() {
  return (
    <>
      <Form method="post">
        <label>
          Insert the token received from your school online catalog
          administration<br></br>
          <input type="text" />
        </label>
        <br />
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
