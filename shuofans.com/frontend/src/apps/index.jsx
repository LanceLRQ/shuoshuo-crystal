import '@styles/index.scss';

import React from 'react';
import {
  Layout, Menu, Row, Col 
} from 'antd';
import styled from 'styled-components';
import ImageLogo from '../images/logo.png';

const LogoBox = styled.div`
  padding: 0 24px;
`;
const LogoImage = styled.img`
  margin-top: -8px;
`;
const LogoTitle = styled.span`
  padding: 16px;
  font-size: 24px;
  color: #1890ff;
  text-shadow: 2px 2px 4px #69c0ff;
  font-family: "QianTuXiaoTu";
`;

const NavigationMenus = [
  { key: 'index', label: '看板' },
  { key: 'voices', label: '声音墙' },
  { key: 'games', label: '小游戏' }
];

function Index() {
  return (
    <Layout className="layout-main-frame">
      <Layout.Header className="layout-main-header">
        <Row>
          <Col xxl={3} xl={4} xs={24} sm={24} lg={6} md={6}>
            <LogoBox>
              <LogoImage alt="logo" width={40} height={40} src={ImageLogo} />
              <LogoTitle className="logo-title">水晶蟹小屋</LogoTitle>
            </LogoBox>
          </Col>
          <Col>
            <Menu
              mode="horizontal"
              defaultSelectedKeys={['index']}
              items={NavigationMenus}
            />
          </Col>
        </Row>
      </Layout.Header>
      <Layout.Content className="layout-main-content">
        Content
      </Layout.Content>
      <div className="layout-main-footer">©2022</div>
    </Layout>
  );
}

export default Index;
