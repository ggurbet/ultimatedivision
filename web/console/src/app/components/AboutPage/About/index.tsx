/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { Switch } from 'react-router-dom';

import { ComponentRoutes, Route } from '@/app/routes';
import { AboutMenu } from '../AboutMenu';

import './index.scss';

const About:React.FC<{children: ComponentRoutes[]}> = ({ children }) =>
    <div className="about">
        <AboutMenu />
        <div className="about__wrapper">
            <Switch>
                {children.map((item: ComponentRoutes, index: number) =>
                    <Route
                        key={index}
                        path={item.path}
                        component={item.component}
                        exact={item.exact}
                    />
                )}
            </Switch>
        </div>
    </div>;


export default About;
