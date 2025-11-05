import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { ExperimentOutlined, DownOutlined, RightOutlined } from '@ant-design/icons';

const SideMenu: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const currentPath = location.pathname;
  const [open, setOpen] = useState(false); // 默认为收起

  const subMenuItems = [
    { key: '/', label: 'AI 模型训练数据' },
    { key: '/model-training', label: 'AI 模型训练' },
    { key: '/data-inference', label: 'AI 模型数据推理' },
    { key: '/fit-model-train-data', label: '机理模型训练数据' },
    { key: '/fit-model-train', label: '机理模型训练' },
    { key: '/fit-model-train-data-inference', label: '机理模型数据推理' },
  ];

  return (
    <div className="menu-container">
      {/* 父级菜单：模型训练 */}
      <div
        className={`menu-item ${currentPath === '/model-training' ? 'active' : ''}`}
        onClick={() => {
          setOpen(prev => !prev);
          // 点击同时跳转到模型训练页面
          navigate('/model-training');
        }}
      >
        <span className="menu-icon"><ExperimentOutlined /></span>
        <span className="menu-text" style={{ whiteSpace: 'nowrap' }}>模型训练</span>
        <span style={{ marginLeft: 'auto', display: 'flex', alignItems: 'center' }}>
          {open ? <DownOutlined /> : <RightOutlined />}
        </span>
      </div>

      {/* 子菜单列表 */}
      {open && (
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
    </div>
  );
};

export default SideMenu;
