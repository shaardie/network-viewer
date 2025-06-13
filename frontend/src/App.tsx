import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Layout } from "./components/Layout";
import { Home } from "./pages/Home";
import { SubnetList } from "./pages/subnet/SubnetList";
import { SubnetCreate } from "./pages/subnet/SubnetCreate";
import { IPList } from "./pages/ip/IPList";

function App() {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/subnet" element={<SubnetList />} />
          <Route path="/subnet/create" element={<SubnetCreate />} />
          <Route path="/ip" element={<IPList />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}

export default App;
