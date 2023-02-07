(require 'ox-publish)

(setq org-publish-project-alist
      (list
       (list "org-site"
             :recursive t
             :base-directory "~/Projects/Org-Site/Development"
             :publishing-directory "~/Projects/Org-Site/Production/Src"
             :publishing-function 'org-twbs-export-to-html)))

(org-publish-all t)

(message "Build Complete")
