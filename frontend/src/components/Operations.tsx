import { useNavigate } from "react-router-dom";

export function Operations({
  id,
  type,
  onDelete,
}: {
  id: number;
  type: string;
  onDelete?: () => void;
}) {
  const navigate = useNavigate();
  const handleDelete = async () => {
    const res = await fetch(`/api/v1/${type}/${id}`, {
      method: "DELETE",
    });

    if (res.ok) {
      if (onDelete) {
        onDelete();
      }
    } else {
      alert("deletion failed");
    }
  };

  return (
    <>
      <button onClick={() => navigate(`/${type}/edit/${id}`)}>Edit</button>
      <button onClick={handleDelete}>Delete</button>
    </>
  );
}
