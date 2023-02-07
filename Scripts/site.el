(require 'ox-publish)

(setq org-publish-project-alist
      (list
       (list "my-org-site"
             :recursive t
             :base-directory "../Development"
             :publishing-directory "../Production/Src"
             :publishing-function 'org-twbs-export-to-html)))

(org-publish-all t)

(message "Build Complete")
