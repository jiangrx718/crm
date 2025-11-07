import React from 'react';
import { Card, Form, Input, Button } from 'antd';
import { useNavigate } from 'react-router-dom';

const Login: React.FC = () => {
  const navigate = useNavigate();
  const [form] = Form.useForm();

  const onFinish = () => {
    // 这里不做登录态校验，提交后直接进入后台首页
    navigate('/home');
  };

  return (
    <div style={{ minHeight: '100vh', display: 'flex', alignItems: 'center', justifyContent: 'center', position: 'relative', overflow: 'hidden' }}>
      {/* 背景：与项目蓝色主色调相匹配的渐变与星空点缀 */}
      <style>
        {`
          .login-bg::before {
            content: '';
            position: absolute;
            inset: 0;
            background: radial-gradient(1000px 600px at 20% 30%, rgba(22,119,255,0.25), transparent),
                        radial-gradient(800px 500px at 80% 70%, rgba(64,169,255,0.25), transparent),
                        linear-gradient(120deg, #0d1b2a 0%, #1b263b 40%, #12253a 100%);
            z-index: -2;
          }
          .login-bg::after {
            content: '';
            position: absolute;
            inset: 0;
            background-image: radial-gradient(rgba(255,255,255,0.12) 1px, transparent 1px);
            background-size: 3px 3px;
            opacity: 0.3;
            z-index: -1;
          }
        `}
      </style>

      <div className="login-bg" style={{ position: 'absolute', inset: 0 }} />

      {/* 登录卡片：左右布局，右侧表单符合 Ant Design 风格 */}
      <Card style={{ width: 780, borderRadius: 12, overflow: 'hidden', padding: 0 }} bodyStyle={{ padding: 0 }}>
        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr' }}>
          {/* 左侧品牌/插图 */}
          <div style={{ background: 'linear-gradient(180deg, #e6f4ff 0%, #f7fbff 100%)', padding: 24, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            <div style={{ maxWidth: 320 }}>
              <div style={{ fontSize: 18, fontWeight: 600, color: '#1677ff', marginBottom: 8 }}>CRM后台管理系统</div>
              <div style={{ color: '#666' }}>自己的，才是最好的。每个企业都应该拥有自己的CRM系统。</div>
              <div style={{ marginTop: 16, height: 160, borderRadius: 8, background: 'url(https://images.unsplash.com/photo-1518770660439-4636190af475?q=80&w=1200&auto=format&fit=crop) center/cover no-repeat' }} />
            </div>
          </div>
          {/* 右侧表单 */}
          <div style={{ padding: 36 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 24 }}>
              <img src="/vite.svg" alt="logo" width={32} height={32} />
              <div style={{ fontSize: 20, fontWeight: 600 }}>CRM商品管理后台系统</div>
            </div>
            <Form form={form} layout="vertical" onFinish={onFinish} requiredMark={false} initialValues={{ account: '' }}>
              <Form.Item label="账号" name="account" rules={[{ required: true, message: '请输入账号' }]}> 
                <Input placeholder="" />
              </Form.Item>
              <Form.Item label="密码" name="password" rules={[{ required: true, message: '请输入密码' }]}> 
                <Input.Password placeholder="" />
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit" block>登录</Button>
              </Form.Item>
            </Form>
          </div>
        </div>
      </Card>
    </div>
  );
};

export default Login;