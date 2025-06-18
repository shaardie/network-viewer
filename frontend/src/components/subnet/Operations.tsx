import { type Subnet } from "../../types/models";
import { useNavigate } from "react-router-dom";

export function Operations({ subnet }: { subnet: Subnet }) {
  const navigate = useNavigate();
  const handleDelete = async () => {
    const res = await fetch(`/api/v1/subnet/${subnet.id}`, {
      method: "DELETE",
    });

    if (res.ok) {
      navigate("/subnet");
    } else {
      alert("deletion failed");
    }
  };

  return (
    <>
      <button onClick={() => navigate(`/subnet/edit/${subnet.id}`)}>
        Edit
      </button>
      <button onClick={handleDelete}>Delete</button>
    </>
  );
}
