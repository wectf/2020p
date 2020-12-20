/* babyrev extension for PHP */

#ifndef PHP_BABYREV_H
# define PHP_BABYREV_H

extern zend_module_entry babyrev_module_entry;
# define phpext_babyrev_ptr &babyrev_module_entry

# define PHP_BABYREV_VERSION "0.1.0"

# if defined(ZTS) && defined(COMPILE_DL_BABYREV)
ZEND_TSRMLS_CACHE_EXTERN()
# endif

#endif	/* PHP_BABYREV_H */

