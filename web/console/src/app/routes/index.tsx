/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { lazy } from 'react';
import { RouteProps, Switch } from 'react-router-dom';

const FootballerCard = lazy(() => import('@components/FootballerCardPage/FootballerCard'));
const FootballField = lazy(() => import('@components/FootballFieldPage/FootballField'));
const MarketPlace = lazy(() => import('@components/MarketPlacePage/MarketPlace'));
const About = lazy(() => import('@components/AboutPage/About'));

import Summary from '@/app/components/AboutPage/WhitePaperPage/Summary';
import GameMechanics from '@/app/components/AboutPage/WhitePaperPage/GameMechanics';
import PayToEarnEconomy from '@components/AboutPage/WhitePaperPage/PayToEarnEconomy';
import Technology from '@components/AboutPage/WhitePaperPage/Technology';

import Fund from '@components/AboutPage/TokenomicsPage/Fund';
import PayToEarn from '@components/AboutPage/TokenomicsPage/PayToEarn';
import Spending from '@components/AboutPage/TokenomicsPage/Spending';
import Staking from '@components/AboutPage/TokenomicsPage/Staking';

/** Route base config implementation */
export class ComponentRoutes {
    /** data route config*/
    constructor(
        public path: string,
        public component: React.FC<{children: ComponentRoutes[]}>,
        public exact: boolean,
        public subRoutes?: ComponentRoutes[]
    ) { }
};

/** Route config implementation */
export class RouteConfig {
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        '/test/marketplace',
        MarketPlace,
        true,
    );
    public static FootballerCard: ComponentRoutes = new ComponentRoutes(
        '/test/marketplace/card',
        FootballerCard,
        true,
    );
    public static FootballField: ComponentRoutes = new ComponentRoutes(
        '/test/field',
        FootballField,
        true,
    );
    public static MyCards: ComponentRoutes = new ComponentRoutes(
        '/test/marketplace/club',
        MarketPlace,
        true,
    );
    public static Tokenomics: ComponentRoutes = new ComponentRoutes(
        '/test/tokenomics',
        About,
        false,
        [
            new ComponentRoutes(
                '/test/tokenomics/udt-spending',
                Spending,
                true
            ),
            new ComponentRoutes(
                '/test/tokenomics/pay-to-earn',
                PayToEarn,
                true
            ),
            new ComponentRoutes(
                '/test/tokenomics/staking',
                Staking,
                true
            ),
            new ComponentRoutes(
                '/test/tokenomics/ud-dao-fund',
                Fund,
                true
            ),
        ]);
    public static Default: ComponentRoutes = new ComponentRoutes(
        '/test/whitepaper',
        About,
        false,
        [
            new ComponentRoutes(
                '/test/whitepaper/',
                Summary,
                true
            ),
            new ComponentRoutes(
                '/test/whitepaper/game-mechanicks',
                GameMechanics,
                true
            ),
            new ComponentRoutes(
                '/test/whitepaper/pay-to-earn-and-economy',
                PayToEarnEconomy,
                true
            ),
            new ComponentRoutes(
                '/test/whitepaper/technology',
                Technology,
                true
            ),
        ]
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.MarketPlace,
        RouteConfig.FootballerCard,
        RouteConfig.FootballField,
        RouteConfig.MyCards,
        RouteConfig.Tokenomics,
        RouteConfig.Default,
    ];
};

// type RoutesProps = { component: React.FC<{ routes?: ComponentRoutes[] }> } & RouteProps;

export const Route: React.FC<ComponentRoutes> = ({
    component: Component,
    ...children
}) => <Component {...children} />;

export const Routes = () =>
    <Switch>
        {RouteConfig.routes.map((route, index) =>
            <Route
                key={index}
                path={route.path}
                component={route.component}
                exact={route.exact}
                children={route.subRoutes}
            />,
        )}
    </Switch>;
