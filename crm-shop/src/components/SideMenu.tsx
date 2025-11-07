import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { ExperimentOutlined, DownOutlined, RightOutlined, SafetyOutlined, ShoppingOutlined, SettingOutlined, HomeOutlined, ShoppingCartOutlined } from '@ant-design/icons';

const SideMenu: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const currentPath = location.pathname;
  const [openTrain, setOpenTrain] = useState(false); // 模型训练：默认收起
  const [openAdmin, setOpenAdmin] = useState(false); // 管理权限：默认收起
  const [openGoods, setOpenGoods] = useState(false); // 商品管理：默认收起
  const [openOrders, setOpenOrders] = useState(false); // 订单管理：默认收起
  const [openSettings, setOpenSettings] = useState(false); // 系统设置：默认收起

  const subMenuItems = [
    { key: '/', label: 'AI 模型训练数据' },
    { key: '/model-training', label: 'AI 模型训练' },
    { key: '/data-inference', label: 'AI 模型数据推理' },
    { key: '/fit-model-train-data', label: '机理模型训练数据' },
    { key: '/fit-model-train', label: '机理模型训练' },
    { key: '/fit-model-train-data-inference', label: '机理模型数据推理' },
  ];

  const adminMenuItems = [
    { key: '/roles', label: '角色管理' },
    { key: '/admins', label: '管理员列表' },
    { key: '/permissions', label: '权限设置' },
  ];

  const goodsMenuItems = [
    { key: '/product-category', label: '商品分类' },
    { key: '/product-list', label: '商品列表' },
  ];

  const ordersMenuItems = [
    { key: '/order-list', label: '订单列表' },
    { key: '/order-statistics', label: '订单统计' },
  ];

  const settingsMenuItems = [
    { key: '/base-settings', label: '基础设置' },
    { key: '/agreement-settings', label: '协议设置' },
  ];

  return (
    <div className="menu-container">
      {/* 顶部菜单：首页 */}
      <div
        className={`menu-item ${currentPath === '/home' ? 'active' : ''}`}
        onClick={() => navigate('/home')}
      >
        <span className="menu-icon"><HomeOutlined /></span>
        <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>首页</span>
      </div>

      {/* 父级菜单：模型训练 */}
      <div
        className={`menu-item ${currentPath === '/model-training' ? 'active' : ''}`}
        onClick={() => {
          setOpenTrain(prev => !prev);
          // 点击同时跳转到模型训练页面
          navigate('/model-training');
        }}
      >
        <span className="menu-icon"><ExperimentOutlined /></span>
        <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>模型训练</span>
        <span style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center' }}>
          {openTrain ? <DownOutlined /> : <RightOutlined />}
        </span>
      </div>

      {/* 子菜单列表 */}
      {openTrain && (
        <div className="submenu-container">
          {subMenuItems.map(item => (
            <div
              key={item.key}
              className={`submenu-item ${currentPath === item.key ? 'active' : ''}`}
              onClick={() => navigate(item.key)}
            >
              <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>{item.label}</span>
            </div>
          ))}
        </div>
      )}

      {/* 父级菜单：管理权限 */}
      <div
        className={`menu-item ${['/roles','/admins','/permissions'].includes(currentPath) ? 'active' : ''}`}
        onClick={() => setOpenAdmin(prev => !prev)}
      >
        <span className="menu-icon"><SafetyOutlined /></span>
        <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>管理权限</span>
        <span style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center' }}>
          {openAdmin ? <DownOutlined /> : <RightOutlined />}
        </span>
      </div>

      {openAdmin && (
        <div className="submenu-container">
          {adminMenuItems.map(item => (
            <div
              key={item.key}
              className={`submenu-item ${currentPath === item.key ? 'active' : ''}`}
              onClick={() => navigate(item.key)}
            >
              <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>{item.label}</span>
            </div>
          ))}
        </div>
      )}

      {/* 父级菜单：商品管理 */}
      <div
        className={`menu-item ${['/product-category','/product-list'].includes(currentPath) ? 'active' : ''}`}
        onClick={() => setOpenGoods(prev => !prev)}
      >
        <span className="menu-icon"><ShoppingOutlined /></span>
        <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>商品管理</span>
        <span style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center' }}>
          {openGoods ? <DownOutlined /> : <RightOutlined />}
        </span>
      </div>

      {openGoods && (
        <div className="submenu-container">
          {goodsMenuItems.map(item => (
            <div
              key={item.key}
              className={`submenu-item ${currentPath === item.key ? 'active' : ''}`}
              onClick={() => navigate(item.key)}
            >
              <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>{item.label}</span>
            </div>
          ))}
        </div>
      )}

      {/* 父级菜单：订单管理 */}
      <div
        className={`menu-item ${['/order-list','/order-statistics'].includes(currentPath) ? 'active' : ''}`}
        onClick={() => setOpenOrders(prev => !prev)}
      >
        <span className="menu-icon"><ShoppingCartOutlined /></span>
        <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>订单管理</span>
        <span style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center' }}>
          {openOrders ? <DownOutlined /> : <RightOutlined />}
        </span>
      </div>

      {openOrders && (
        <div className="submenu-container">
          {ordersMenuItems.map(item => (
            <div
              key={item.key}
              className={`submenu-item ${currentPath === item.key ? 'active' : ''}`}
              onClick={() => navigate(item.key)}
            >
              <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>{item.label}</span>
            </div>
          ))}
        </div>
      )}

      {/* 父级菜单：系统设置 */}
      <div
        className={`menu-item ${['/base-settings','/agreement-settings'].includes(currentPath) ? 'active' : ''}`}
        onClick={() => setOpenSettings(prev => !prev)}
      >
        <span className="menu-icon"><SettingOutlined /></span>
        <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>系统设置</span>
        <span style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center' }}>
          {openSettings ? <DownOutlined /> : <RightOutlined />}
        </span>
      </div>

      {openSettings && (
        <div className="submenu-container">
          {settingsMenuItems.map(item => (
            <div
              key={item.key}
              className={`submenu-item ${currentPath === item.key ? 'active' : ''}`}
              onClick={() => navigate(item.key)}
            >
              <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>{item.label}</span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default SideMenu;
