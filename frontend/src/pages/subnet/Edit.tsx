import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { type Subnet } from "../../types/models.ts";
import { Form } from "../../components/subnet/Form.tsx";

export function Edit() {
  const { id } = useParams<{ id: string }>();
  const [subnet, setSubnet] = useState<Subnet | null>(null);
  const [error, setError] = useState("");

  useEffect(() => {
    fetch(`/api/v1/subnet/${id}`)
      .then((res) => {
        if (!res.ok) throw new Error("Error while loading");
        return res.json();
      })
      .then(setSubnet)
      .catch((err) => setError(err.message));
  }, [id]);

  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (!subnet) return <p>loading...</p>;
  return (
    <>
      <h1>Edit Subnet</h1>
      <Form subnet={subnet}></Form>
    </>
  );
}
