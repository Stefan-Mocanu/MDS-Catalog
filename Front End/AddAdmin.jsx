import { Form } from "react-router-dom";

export async function action({ request }) {
  const formData = await request.formData();
  // const obj = Object.fromEntries(formData);
  const url = "/api/inserareScoala";
  await fetch(url, {
    method: "POST",
    body: formData,
  })
    .then((response) => {
      if (response.ok) alert("Adaugare rol cu succes!");
    })
    .catch((error) => {});
  return 0;
}

export default function AddAdmin() {
  return (
    <>
      <Form method="post">
        <label>
          Numele scolii
          <input type="text" name="nume" placeholder="Numele scolii" />
        </label>
        <br />
        <button type="submit">Submit</button>
      </Form>
    </>
  );
}
