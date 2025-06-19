import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Layout } from "./components/Layout";
import { Home } from "./pages/Home";
import { List as SubnetList } from "./pages/subnet/List";
import { List as IPList } from "./pages/ip/List";
import { Show as SubnetShow } from "./pages/subnet/Show";
import { Create as SubnetCreate } from "./pages/subnet/Create";
import { Edit as SubnetEdit } from "./pages/subnet/Edit";

export default function App() {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/subnet" element={<SubnetList />} />
          <Route path="/subnet/:id" element={<SubnetShow />} />
          <Route path="/subnet/create" element={<SubnetCreate />} />
          <Route path="/subnet/edit/:id" element={<SubnetEdit />} />
          <Route path="/ip" element={<IPList />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}
