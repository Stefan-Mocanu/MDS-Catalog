import { Form } from "react-router-dom";
action de facut
export default function AddAdmin() {
  return (
    <>
      <Form method="post">
        <input type="text" name="name" placeholder="Name" />
        <input type="submit" name="addadmin" />
      </Form>
    </>
  );
}
