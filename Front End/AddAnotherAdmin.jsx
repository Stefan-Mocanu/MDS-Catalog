import { Form, useOutletContext } from "react-router-dom";

export async function action({ request }) {
  const formData = await request.formData();
  const url = "/api/adaugaAdmin";
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => response.json())
    .then((response) => console.log(response))
    .catch((error) => console.log(error));
  return 0;
}

export default function AddAnotherAdmin() {
  const data = useOutletContext();
  return (
    <>
      <Form method="post">
        <label>
          Insert the user's ID which you want to make admin
          <br />
          <input type="number" name="id_cont" />
        </label>
        <input type="hidden" name="id_scoala" value={data["id"]} />
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
