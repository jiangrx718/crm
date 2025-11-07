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
import ProductCategory from './pages/ProductCategory';
import ProductList from './pages/ProductList';
import BaseSettings from './pages/BaseSettings';
import AgreementSettings from './pages/AgreementSettings';
import Home from './pages/Home';
import OrderList from './pages/OrderList';
import OrderStatistics from './pages/OrderStatistics';
import ArticleCategory from './pages/ArticleCategory';
import ArticleList from './pages/ArticleList';

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
                <Route path="/home" element={<Home />} />
                <Route path="/" element={<DataConversion />} />
                <Route path="/model-training" element={<ModelTraining />} />
                <Route path="/data-inference" element={<DataInference />} />
                <Route path="/fit-model-train-data" element={<FitModelTrainData />} />
                <Route path="/fit-model-train" element={<FitModelTrain />} />
                <Route path="/fit-model-train-data-inference" element={<FitModelTrainDataInference />} />
                <Route path="/roles" element={<RoleManagement />} />
                <Route path="/admins" element={<AdminList />} />
                <Route path="/permissions" element={<PermissionSettings />} />
                <Route path="/product-category" element={<ProductCategory />} />
                <Route path="/product-list" element={<ProductList />} />
                <Route path="/article-category" element={<ArticleCategory />} />
                <Route path="/article-list" element={<ArticleList />} />
                <Route path="/order-list" element={<OrderList />} />
                <Route path="/order-statistics" element={<OrderStatistics />} />
                <Route path="/base-settings" element={<BaseSettings />} />
                <Route path="/agreement-settings" element={<AgreementSettings />} />
              </Routes>
            </Content>
          </Layout>
        </Layout>
      </Layout>
    </Router>
  );
}

export default App;