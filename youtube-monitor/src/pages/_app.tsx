import { Global } from '@emotion/react'
import { appWithTranslation } from 'next-i18next'
import type { AppProps } from 'next/app'
import { globalStyle } from '../styles/global.styles'
import SessionProvider from '../components/SessionProvider'

function App({ Component, pageProps }: AppProps): JSX.Element {
	return (
		<SessionProvider session={pageProps.session}>
			<Global styles={globalStyle} />
			<Component {...pageProps} />
		</SessionProvider>
	)
}

export default appWithTranslation(App)
