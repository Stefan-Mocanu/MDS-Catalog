import { Form } from "react-router-dom";

export async function action({ request }) {
  const formData = await request.formData();
  // const obj = Object.fromEntries(formData);
  const url = "/api/alaturare";
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => response.json())
    .then((response) => console.log(response))
    .catch((error) => console.log(error));
  return 0;
}

export default function AddUser() {
  return (
    <>
      <Form method="post">
        <label>
          Alegeti rolul dorit
          <br />
          <select name="rol">
            <option value="elev">Elev</option>
            <option value="profesor">Profesor</option>
            <option value="parinte">Parinte</option>
          </select>
        </label>
        <br />
        <label>
          Insert the token received from your school online catalog
          administration<br></br>
          <input type="text" name="token" />
        </label>
        <br />
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
