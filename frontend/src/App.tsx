import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Layout } from "./components/Layout";
import { Home } from "./pages/Home";
import { List } from "./pages/subnet/List";
import { IPList } from "./pages/ip/IPList";
import { Show } from "./pages/subnet/Show";
import { Create } from "./pages/subnet/Create";
import { Edit } from "./pages/subnet/Edit";

function App() {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/subnet" element={<List />} />
          <Route path="/subnet/:id" element={<Show />} />
          <Route path="/subnet/create" element={<Create />} />
          <Route path="/subnet/edit/:id" element={<Edit />} />
          <Route path="/ip" element={<IPList />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}

export default App;
