#  ----------------------------------------------------------------
# Copyright 2016 Cisco Systems
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ------------------------------------------------------------------

"""
printer_factory.py

 Returns printer

"""
from ydkgen.printer.cpp.cpp_bindings_printer import CppBindingsPrinter
from ydkgen.printer.python.python_bindings_printer import PythonBindingsPrinter
from ydkgen.printer.go.go_bindings_printer import GoBindingsPrinter

class PrinterFactory(object):

    def get_printer(self, language):
        if language == 'cpp':
            return CppBindingsPrinter
        elif language == 'python':
            return PythonBindingsPrinter
        elif language == 'go':
            return GoBindingsPrinter
        else:
            raise Exception('Language {0} not yet supported'.format(language))
