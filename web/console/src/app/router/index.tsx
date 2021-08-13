// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { lazy } from 'react';
import { Switch } from 'react-router-dom';

const FootballerCard = lazy(() => import('@components/FootballerCardPage/FootballerCard'));
const FootballField = lazy(() => import('@components/FootballFieldPage/FootballField'));
const MarketPlace = lazy(() => import('@components/MarketPlacePage/MarketPlace'));
const WhitePaper = lazy(() => import('@components/AboutPage/WhitePaperPage/WhitePaper'));
const Tokenomics = lazy(() => import('@components/AboutPage/TokenomicsPage/Tokenomics'));

import Summary from '@/app/components/AboutPage/WhitePaperPage/Summary';
import GameMechanics from '@/app/components/AboutPage/WhitePaperPage/GameMechanics';
import PayToEarnEconomy from '@components/AboutPage/WhitePaperPage/PayToEarnEconomy';
import Technology from '@components/AboutPage/WhitePaperPage/Technology';
import Fund from '@components/AboutPage/TokenomicsPage/Fund';
import PlayToEarn from '@components/AboutPage/TokenomicsPage/PlayToEarn';
import Spending from '@components/AboutPage/TokenomicsPage/Spending';
import Staking from '@components/AboutPage/TokenomicsPage/Staking';

/** Route base config implementation */
export class ComponentRoutes {
    /** data route config*/
    constructor(
        public path: string,
        public component: React.FC<any>,
        public exact: boolean,
        public children?: ComponentRoutes[]
    ) { }
    /** Method for creating child subroutes path */
    public with(child: ComponentRoutes, parrent: ComponentRoutes): ComponentRoutes {
        child.path = `${parrent.path}/${child.path}`;

        return this;
    }
    /** Call with method for each child */
    public addChildren(children: ComponentRoutes[]): ComponentRoutes {
        this.children = children.map(item => item.with(item, this));

        return this;
    }
};

/** interfafe fot AboutPage subroutes */
interface RouteItem {
    path: string;
    component: React.FC<any>;
    exact: boolean;
    children?: ComponentRoutes[];
    with?: (child: ComponentRoutes, parrent: ComponentRoutes) => ComponentRoutes;
    addChildren?: (children: ComponentRoutes[]) => ComponentRoutes;
}

/** Route config implementation */
export class RouteConfig {
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        '/marketplace',
        MarketPlace,
        true,
    );
    public static FootballerCard: ComponentRoutes = new ComponentRoutes(
        '/card',
        FootballerCard,
        true,
    );
    public static FootballField: ComponentRoutes = new ComponentRoutes(
        '/field',
        FootballField,
        true,
    );
    public static MyCards: ComponentRoutes = new ComponentRoutes(
        '/club',
        MarketPlace,
        true,
    );
    public static Whitepaper: ComponentRoutes = new ComponentRoutes(
        '/whitepaper',
        WhitePaper,
        false
    );
    public static Tokenomics: ComponentRoutes = new ComponentRoutes(
        '/tokenomics',
        Tokenomics,
        false
    );
    public static Summary: ComponentRoutes = new ComponentRoutes(
        'summary',
        Summary,
        true
    );
    public static GameMechanick: ComponentRoutes = new ComponentRoutes(
        'game-mechanicks',
        GameMechanics,
        true
    );
    public static PayToEarnEconomy: ComponentRoutes = new ComponentRoutes(
        'pay-to-earn-and-economy',
        PayToEarnEconomy,
        true
    );
    public static Technology: ComponentRoutes = new ComponentRoutes(
        'technology',
        Technology,
        true
    );
    public static Spending: ComponentRoutes = new ComponentRoutes(
        'udt-spending',
        Spending,
        true
    );
    public static PayToEarn: ComponentRoutes = new ComponentRoutes(
        'pay-to-earn',
        PlayToEarn,
        true
    );
    public static Staking: ComponentRoutes = new ComponentRoutes(
        'staking',
        Staking,
        true
    );
    public static Fund: ComponentRoutes = new ComponentRoutes(
        'ud-dao-fund',
        Fund,
        true
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        '/',
        MarketPlace,
        true,
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.Default,
        RouteConfig.FootballField,
        RouteConfig.MarketPlace,
        RouteConfig.FootballerCard,
        RouteConfig.MyCards,
        RouteConfig.Whitepaper.addChildren([
            RouteConfig.Summary,
            RouteConfig.GameMechanick,
            RouteConfig.PayToEarnEconomy,
            RouteConfig.Technology,
        ]),
        RouteConfig.Tokenomics.addChildren([
            RouteConfig.Spending,
            RouteConfig.PayToEarn,
            RouteConfig.Staking,
            RouteConfig.Fund,
        ]),
    ];
};

export const Route: React.FC<RouteItem> = ({
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
                children={route.children}
            />
        )}
    </Switch>;
