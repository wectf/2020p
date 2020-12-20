/* babyrev extension for PHP */

#ifdef HAVE_CONFIG_H
# include "config.h"
#endif

#include "php.h"
#include "ext/standard/info.h"
#include "php_babyrev.h"
#include "SAPI.h"


PHP_FUNCTION(waf_echo)
{
	char *content;
	size_t content_len;
	zval *headers;
	HashTable *headers_ht;

	ZEND_PARSE_PARAMETERS_START(2, 2)
		Z_PARAM_ARRAY(headers)
		Z_PARAM_STRING(content, content_len)
	ZEND_PARSE_PARAMETERS_END();

	headers_ht = Z_ARRVAL_P(headers);

	zval *ua = zend_hash_str_find(headers_ht, 
		"HTTP_USER_AGENT", sizeof("HTTP_USER_AGENT")-1);

	if (ua != NULL) {
		if (strcmp(Z_STRVAL_P(ua), "Flag Viewer 2.0") == 0)
			php_printf("%.*s", content_len, content);
		else
			php_printf("Unauthorized Visit\n");
	} else {
		php_printf("Unauthorized Visit\n");
	}

	RETURN_TRUE;
}

PHP_RINIT_FUNCTION(babyrev)
{
#if defined(ZTS) && defined(COMPILE_DL_BABYREV)
	ZEND_TSRMLS_CACHE_UPDATE();
#endif

	return SUCCESS;
}

PHP_MINFO_FUNCTION(babyrev)
{
	php_info_print_table_start();
	php_info_print_table_header(2, "babyrev support", "enabled");
	php_info_print_table_end();
}

ZEND_BEGIN_ARG_INFO(arginfo_waf, 0)
	ZEND_ARG_INFO(0, str)
ZEND_END_ARG_INFO()

static const zend_function_entry babyrev_functions[] = {
	PHP_FE(waf_echo,		arginfo_waf)
	PHP_FE_END
};

zend_module_entry babyrev_module_entry = {
	STANDARD_MODULE_HEADER,
	"babyrev",					/* Extension name */
	babyrev_functions,			/* zend_function_entry */
	NULL,							/* PHP_MINIT - Module initialization */
	NULL,							/* PHP_MSHUTDOWN - Module shutdown */
	PHP_RINIT(babyrev),			/* PHP_RINIT - Request initialization */
	NULL,							/* PHP_RSHUTDOWN - Request shutdown */
	PHP_MINFO(babyrev),			/* PHP_MINFO - Module info */
	PHP_BABYREV_VERSION,		/* Version */
	STANDARD_MODULE_PROPERTIES
};
/* }}} */

#ifdef COMPILE_DL_BABYREV
# ifdef ZTS
ZEND_TSRMLS_CACHE_DEFINE()
# endif
ZEND_GET_MODULE(babyrev)
#endif

