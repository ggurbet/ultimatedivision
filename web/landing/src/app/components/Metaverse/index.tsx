// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import CardCenter from '@static/images/metaverse/card-center.png';
import CardCenterMedium from '@static/images/metaverse/card-center-834.png';
import CardCenterSmall from '@static/images/metaverse/card-center-414.png';
import CardLeft from '@static/images/metaverse/card-left.png';
import CardLeftMedium from '@static/images/metaverse/card-left-834.png';
import CardLeftSmall from '@static/images/metaverse/card-left-414.png';
import CardRight from '@static/images/metaverse/card-right.png';
import CardRightMedium from '@static/images/metaverse/card-right-834.png';
import CardRightSmall from '@static/images/metaverse/card-right-414.png';

import './index.scss';

export const Metaverse: React.FC = () => {
    return (
        <section className="metaverse">
            <div className="wrapper">
                <span className="metaverse__title">Ultimate Divison</span>
                <span className="metaverse__subtitle">Football Metaverse</span>
                <div className="metaverse__cards">
                    <picture className="left-card">
                        <source media="(max-width: 414px)" srcSet={CardLeftSmall} />
                        <source media="(max-width: 834px)" srcSet={CardLeftMedium} />
                        <source media="(min-width: 1440px)" srcSet={CardLeft}/>
                        <img src={CardLeft} alt="Left player"></img>
                    </picture>
                    <picture className="center-card">
                        <source media="(max-width: 414px)" srcSet={CardCenterSmall} />
                        <source media="(max-width: 834px)" srcSet={CardCenterMedium} />
                        <source media="(min-width: 1440px)" srcSet={CardCenter}/>
                        <img src={CardCenter} alt="Main player"></img>
                    </picture>
                    <picture className="right-card">
                        <source media="(max-width: 414px)" srcSet={CardRightSmall} />
                        <source media="(max-width: 834px)" srcSet={CardRightMedium} />
                        <source media="(min-width: 1440px)" srcSet={CardRight}/>    
                        <img src={CardRight} alt="Right player"></img>
                    </picture>
                </div>
                <div className="metaverse__sold-scale">
                    <span className="metaverse__sold-scale-text">Cards Sold 0/10000</span>
                </div>
                {/**TODO: Need merge with PR #194 */}
                {/* <MintButton text="Mint"/> */}
            </div>
        </section>
    );
};
