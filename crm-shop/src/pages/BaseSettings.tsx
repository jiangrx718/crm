import React from 'react';
import { Breadcrumb, Card, Form, Input, InputNumber, Switch, Button, Space } from 'antd';

const BaseSettings: React.FC = () => {
  const [form] = Form.useForm();

  const onSubmit = () => {
    form.validateFields().then(() => {
      // 这里可集成保存接口
    });
  };

  return (
    <div className="page-container">
      <Breadcrumb items={[{ title: '首页' }, { title: '系统设置' }, { title: '基础设置' }]} />

      <Card style={{ marginTop: 16 }}>
        <Form
          form={form}
          layout="vertical"
          requiredMark={false}
          initialValues={{ siteName: '', hotline: '', pageSize: 10, enableRegister: true }}
        >
          <Space direction="vertical" style={{ width: '100%' }} size={16}>
            <Form.Item label="网站名称" name="siteName" rules={[{ required: true, message: '请输入网站名称' }]}> 
              <Input placeholder="请输入网站名称" />
            </Form.Item>

            <Form.Item label="LOGO地址" name="logoUrl" rules={[{ type: 'url', message: '请输入合法的URL' }]}> 
              <Input placeholder="https://example.com/logo.png" />
            </Form.Item>

            <Form.Item label="客服热线" name="hotline"> 
              <Input placeholder="400-000-0000" />
            </Form.Item>

            <Form.Item label="分页大小" name="pageSize" rules={[{ required: true, message: '请输入分页大小' }]}> 
              <InputNumber style={{ width: 160 }} min={5} max={100} />
            </Form.Item>

            <Form.Item label="是否开启注册" name="enableRegister" valuePropName="checked"> 
              <Switch />
            </Form.Item>

            <Form.Item>
              <Button type="primary" onClick={onSubmit}>
                保存
              </Button>
            </Form.Item>
          </Space>
        </Form>
      </Card>
    </div>
  );
};

export default BaseSettings;