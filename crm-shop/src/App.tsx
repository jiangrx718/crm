import { Layout } from 'antd';
import { HashRouter as Router, Route, Routes } from 'react-router-dom'; // 修改这里：BrowserRouter → HashRouter
import SideMenu from './components/SideMenu';
import DataConversion from './pages/DataConversion';
import ModelTraining from './pages/ModelTraining';
import DataInference from './pages/DataInference';
import RoleManagement from './pages/RoleManagement';
import AdminList from './pages/AdminList';
import PermissionSettings from './pages/PermissionSettings';
import FitModelTrainData from './pages/FitModelTrainData';
import FitModelTrain from './pages/FitModelTrain';
import FitModelTrainDataInference from './pages/FitModelTrainDataInference';

const { Header, Content, Sider } = Layout;

function App() {
  return (
    <Router>
      <Layout className="app-container">
        <Header className="app-header">
          <div className="app-logo">CRM商品管理后台系统</div>
        </Header>
        <Layout className="app-layout">
          <Sider className="app-sider" width={200}>
            <SideMenu />
          </Sider>
          <Layout>
            <Content className="app-content">
              <Routes>
                <Route path="/" element={<DataConversion />} />
                <Route path="/model-training" element={<ModelTraining />} />
                <Route path="/data-inference" element={<DataInference />} />
                <Route path="/fit-model-train-data" element={<FitModelTrainData />} />
                <Route path="/fit-model-train" element={<FitModelTrain />} />
                <Route path="/fit-model-train-data-inference" element={<FitModelTrainDataInference />} />
                <Route path="/roles" element={<RoleManagement />} />
                <Route path="/admins" element={<AdminList />} />
                <Route path="/permissions" element={<PermissionSettings />} />
              </Routes>
            </Content>
          </Layout>
        </Layout>
      </Layout>
    </Router>
  );
}

export default App;